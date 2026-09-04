package providers

import (
	"strings"
)

// valida se o formato só usa placeholders conhecidos e não deixa chave órfã (ex: "{user" sem fechar)
func validateProxyUrlPattern(urlPattern string) bool {
	matches := proxyUrlPlaceholderPattern.FindAllString(urlPattern, -1)
	if len(matches) == 0 {
		return false
	}

	for _, placeholder := range matches {
		if !allowedProxyUrlPlaceholders[placeholder] {
			return false
		}
	}

	stripped := proxyUrlPlaceholderPattern.ReplaceAllString(urlPattern, "")
	return !strings.ContainsAny(stripped, "{}")
}
