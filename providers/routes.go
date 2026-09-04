package providers

import (
	"net/http"
	"talos/shared"
)

type httpStatus int

var AddClientOk = "Cliente Adicionado"
var AddClientNok = "Erro em persistir no banco"

// AddClientPersistence is assigned by the application layer.
var AddClientPersistence func(Client) bool

func AddClient(client Client, urlPattern string) shared.Response {
	// se for preenchido, validar
	// se não for preenchido, preencher
	if client.ProxyUrl == nil || !validateClientProxyUrlPattern(client.ProxyUrl) {
		url, errMsg := getProxyUrl(client.Proxy, client.Proxy.UrlPattern)
		if errMsg != "" {
			_ = url
			return shared.Response{Status: http.StatusBadRequest, Msg: errMsg}
		}
		client.ProxyUrl = &url
	}

	res := AddClientPersistence != nil && AddClientPersistence(client)
	if !res {
		return shared.Response{Status: http.StatusInternalServerError, Msg: AddClientNok}
	}

	return shared.Response{Status: http.StatusOK, Msg: AddClientOk}
}
