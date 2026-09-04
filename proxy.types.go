package main

import "regexp"

type ClientProxy struct {
	Id    string `json:"id"`
	Proxy Proxy  `json:"Proxy"`
}

type Proxy struct {
	User string `json:"user"`
	Pass string `json:"pass"`
	Host string `json:"host"`
	Port string `json:"Port"`
	Ttl  string `json:"ttl"`
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
