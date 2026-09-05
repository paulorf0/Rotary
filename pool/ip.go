package pool

import "time"

type (
	State int
)

var (
	Available State = 0
	Leased    State = 1
	Burned    State = 2
)

type IP struct {
	ID         string
	ProxyURL   string
	LastReport Report // QUESTION: É necessário guardar o último report ou existe forma mais simples?
	// Se for a primeira vez, não tem report anterior (nil)
	Score   float32   // TODO: calcular como score = successRate*w1 + latencyScore*w2 (rolling), não decremento fixo
	ScoreAt time.Time // quando foi materializado
	State   State     // available | leased | burned
	// TODO: cooldown/backoff antes de reoferecer o ip após report ruim, usando LastUsedAt/ScoreAt
	// (backoff exponencial se reincide, não liberar na hora)
	LastUsedAt time.Time // desempate LRU
	Lease      Lease
}

func (i *IP) ReportCount() {
}

func (i *IP) LeaseIp(clientID string) {
	now := time.Now()

	i.State = Leased
	i.Lease.ClosedAt = nil
	i.Lease.AcquiredAt = &now
	i.Lease.ClientID = &clientID
}

func (i *IP) BurnIP() {
	// O gerenciador vai passar por todos os ips e verificar qual está burned. Aquele que estiver, um novo ip vai ser gerado para substituir.
	// Nenhum IP do pool é apagado, se não teria que ficar recriando structs.

	now := time.Now()

	i.State = Burned
	i.Lease.ClosedAt = &now
	i.Lease.AcquiredAt = nil
	i.Lease.ClientID = nil
}

func (i *IP) FreeIP() {
	now := time.Now()

	i.State = Available
	i.Lease.ClosedAt = &now
	i.Lease.AcquiredAt = nil
	i.Lease.ClientID = nil
}

func (i *IP) GetScore() float32 {
	// TODO: Como escolher os pesos?
	var w1 float32 = 1.0
	var w2 float32 = 1.0

	return i.Score*w1 + i.LastReport.RequestTime*w2
}
