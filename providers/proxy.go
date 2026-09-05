package providers

import (
	"strings"
)

func getProxyURL(proxyStruct Proxy, urlPattern string) (string, string) {
	if !validateProxyURLPattern(urlPattern) {
		return "", ErrProxyPatternNotValid
	}

	proxyReplacer := strings.NewReplacer(
		"{user}", proxyStruct.User,
		"{pass}", proxyStruct.Pass,
		"{host}", proxyStruct.Host,
		"{port}", proxyStruct.Port,
		"{ttl}", proxyStruct.TTL,
	)

	return proxyReplacer.Replace(urlPattern), ""
}
