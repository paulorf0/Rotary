package main

import (
	"strings"
)

func getProxyUrl(proxyStruct Proxy, urlPattern string) (ProxyUrl, error) {
	if !validateProxyUrlPattern(urlPattern) {
		return "", ErrProxyPatternNotValid
	}

	var proxyReplacer = strings.NewReplacer(
		"{user}", proxyStruct.User,
		"{pass}", proxyStruct.Pass,
		"{host}", proxyStruct.Host,
		"{port}", proxyStruct.Port,
		"{ttl}", proxyStruct.Ttl,
	)

	return ProxyUrl(proxyReplacer.Replace(urlPattern)), nil
}
