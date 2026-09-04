package providers

import (
	"strings"
)

func getProxyUrl(proxyStruct Proxy, urlPattern string) (ProxyUrl, string) {
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

	return ProxyUrl(proxyReplacer.Replace(urlPattern)), ""
}
