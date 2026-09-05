package providers

import "regexp"

// TODO: Talvez isso seja lógica para ficar no DB. É a representação do cliente no banco de dados.
// As informações do cliente nunca vão ser puxadas todas para a memória.
type Client struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Proxy Proxy  `json:"proxy"`
	// ProxyURL é saída: montado a partir de Proxy.UrlPattern e dos campos de Proxy.
	// O que o cliente enviar nesse campo é descartado.
	ProxyURL *string `json:"proxyUrl"`
}

type Proxy struct {
	User       string `json:"user"`
	Pass       string `json:"pass"`
	Host       string `json:"host"`
	Port       string `json:"port"`
	TTL        string `json:"ttl"`
	URLPattern string `json:"urlPattern"`
}

var (
	proxyURLPlaceholderPattern  = regexp.MustCompile(`\{[^{}]*\}`)
	allowedProxyURLPlaceholders = map[string]bool{
		"{user}": true,
		"{pass}": true,
		"{host}": true,
		"{port}": true,
		"{ttl}":  true,
	}
)
