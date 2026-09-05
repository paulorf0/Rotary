package providers

import (
	"net/http"

	"talos/shared"
)

var (
	AddClientOk  = "Cliente Adicionado"
	AddClientNok = "Erro em persistir no banco"
)

// AddClientPersistence is assigned by the application layer.
var AddClientPersistence func(Client) bool

func AddClient(client Client) shared.Response {
	resolvedURL, errMsg := getProxyURL(client.Proxy, client.Proxy.URLPattern)
	if errMsg != "" {
		return shared.Response{Status: http.StatusBadRequest, Msg: errMsg}
	}
	client.ProxyURL = &resolvedURL

	res := AddClientPersistence != nil && AddClientPersistence(client)
	if !res {
		return shared.Response{Status: http.StatusInternalServerError, Msg: AddClientNok}
	}

	return shared.Response{Status: http.StatusOK, Msg: AddClientOk}
}
