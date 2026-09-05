package pool

import "time"

// TODO: refatorar (nomes, tipo, organização). Por enquanto só pra mapear os status que a rotação precisa tratar.
//
// Sucesso -> reforça score, ip continua bom.
// StatusOK = 200

// Bloqueio/ban pelo antibot -> decai score rápido, pode ir direto pra Burned.
// StatusForbidden = 403 // ban do antibot no alvo (Cloudflare, PerimeterX, etc)
// StatusUnauthorized = 401 // às vezes usado como "challenge wall" por antibot, não auth real

// Rate limit -> temporário, cooldown/backoff no ip, não burn.
// StatusTooManyRequests = 429

// Falha do proxy em si (não do alvo) -> problema de infra do proxy, tratar diferente de ban.
// StatusProxyAuthRequired = 407 // proxy upstream rejeitou credencial
// StatusBadGateway = 502 // proxy não conseguiu falar com o alvo
// StatusGatewayTimeout = 504 // proxy conectou mas alvo não respondeu a tempo

// Erro do alvo, não culpa do ip -> não devia penalizar o proxy.
// StatusInternalServerError = 500
// StatusServiceUnavailable = 503

// Timeout/conexão -> ambíguo (pode ser proxy morto ou alvo lento). Sem status HTTP (falha de conexão/DNS/timeout de socket) também cai aqui.
// StatusRequestTimeout = 408

// Challenge/CAPTCHA -> muitas vezes vem como 200 com corpo de desafio (JS challenge), não dá pra confiar só no status code.
// Precisa inspecionar corpo/headers (ex: presença de cf-mitigated, texto de captcha) se quiser detectar isso.

func (p *pool) poolBelongsClient(clientID string) bool {
	return p.ClientID == clientID
}

// TEST: Acredito que a lógica não esteja funcionando tão bem.
func (p *pool) getIP(clientID string) (string, error) {
	if !p.poolBelongsClient(clientID) {
		return "", NewError(PoolDoesNotBelongClient)
	}

	ip, err := p.getIPByClientID(clientID)
	if err != nil {
		return "", err
	}

	ip.LastUsedAt = time.Now()

	return ip.ProxyURL, nil
}

func (p *pool) getIPByClientID(clientID string) (*IP, error) {
	for i := range p.Ips {
		if p.Ips[i].State == Leased && *p.Ips[i].Lease.ClientID == clientID {
			return &p.Ips[i], nil
		}
	}
	return nil, NewError(NoIPAssignedToTheClient)
}

// O cliente vai fazer a requisição para o target.
// Na resposta, ele vai reportar para o backend as informações que foi retornada (struct Report)
// Nesse ponto, o backend precisa decidir se vai trocar o ip do cliente, para a próxima vez que ele pegar o ip, rotacionar ou não.
// TODO: decisão de rotacionar deve considerar:
// - categoria do HTTPStatus (ban/rate-limit/proxy-fail/target-error/ambíguo, ver topo do arquivo)
// - latencyScore do report (recalcular Score = successRate*w1 + latencyScore*w2)
// - velocity: quantos reports ruins numa janela de tempo, não total acumulado
// - threshold de Score pra virar Burned (BurnedThreshold)
// - cooldown/backoff antes de reoferecer o ip, não liberar na hora
// - se o lease do cliente ainda tá aberto (rotação é fim-de-lease, não por-request)
func (p *pool) report(clientID string, report Report) error {
	if !p.poolBelongsClient(clientID) {
		return NewError(PoolDoesNotBelongClient)
	}

	ip, err := p.getIPByClientID(clientID)
	// TODO: tratar erro?
	if err != nil {
		return err
	}
	// Garantido que o cliente que reportou tem um ip leased e o pool pertence ao cliente.
	ip.LastReport = report // QUESTION: Precisa da propriedade LastReport?

	newScore := ip.GetScore()
	ip.Score = newScore

	// O usuário não quer rotacionar o ip, portanto, apenas armazena as estatisticas e continua.
	if !report.Rotate {
		return nil
	}

	// Lógica de rotacionamento de ip

	// Se ip queimado, já vai ter saído da função. Se não saiu anida, verifica o ttl e rotaciona caso expirado.
	if time.Since(*ip.Lease.AcquiredAt).Seconds() > ip.Lease.ExpiresAt {
		// O tempo que foi adquirido passou o TTL.
		ip.FreeIP()

		// TODO:  Rotacionar o ip.
	}

	return NewError(ReportUnsucessful)
}
