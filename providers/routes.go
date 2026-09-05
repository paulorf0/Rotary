package providers

import (
	"net/http"
	"talos/shared"
)

var AddClientOk = "Cliente Adicionado"
var AddClientNok = "Erro em persistir no banco"

// AddClientPersistence is assigned by the application layer.
var AddClientPersistence func(Client) bool

func AddClient(client Client) shared.Response {
	resolvedUrl, errMsg := getProxyUrl(client.Proxy, client.Proxy.UrlPattern)
	if errMsg != "" {
		return shared.Response{Status: http.StatusBadRequest, Msg: errMsg}
	}
	client.ProxyUrl = &resolvedUrl

	res := AddClientPersistence != nil && AddClientPersistence(client)
	if !res {
		return shared.Response{Status: http.StatusInternalServerError, Msg: AddClientNok}
	}

	return shared.Response{Status: http.StatusOK, Msg: AddClientOk}
}
