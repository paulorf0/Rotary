package providers

import (
	"strings"
)

// valida se o formato só usa placeholders conhecidos e não deixa chave órfã (ex: "{user" sem fechar)
func validateProxyURLPattern(urlPattern string) bool {
	matches := proxyURLPlaceholderPattern.FindAllString(urlPattern, -1)
	if len(matches) == 0 {
		return false
	}

	for _, placeholder := range matches {
		if !allowedProxyURLPlaceholders[placeholder] {
			return false
		}
	}

	stripped := proxyURLPlaceholderPattern.ReplaceAllString(urlPattern, "")
	return !strings.ContainsAny(stripped, "{}")
}
