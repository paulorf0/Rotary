package providers

import "regexp"

type Client struct {
	Id       string  `json: "id"`
	Name     string  `json: "name"`
	Proxy    Proxy   `json: "proxy"`
	ProxyUrl *string `json: "proxyUrl`
}

type Proxy struct {
	User       string `json: "user"`
	Pass       string `json: "pass"`
	Host       string `json: "host"`
	Port       string `json: "port"`
	Ttl        string `json: "ttl"`
	urlPattern string `json: "urlPattern"`
}

type ProxyUrl string

var proxyUrlPlaceholderPattern = regexp.MustCompile(`\{[^{}]*\}`)
var allowedProxyUrlPlaceholders = map[string]bool{
	"{user}": true,
	"{pass}": true,
	"{host}": true,
	"{port}": true,
	"{ttl}":  true,
}
