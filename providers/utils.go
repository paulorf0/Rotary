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

// valida o ProxyUrl vindo do cliente (campo ponteiro, pode não ter sido enviado) — nil conta como inválido
func validateClientProxyUrlPattern(urlPattern *string) bool {
	if urlPattern == nil {
		return false
	}
	return validateProxyUrlPattern(*urlPattern)
}
