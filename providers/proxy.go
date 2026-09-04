package providers

import (
	"strings"
)

func getProxyUrl(proxyStruct Proxy, urlPattern string) (string, string) {
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

	return proxyReplacer.Replace(urlPattern), ""
}
