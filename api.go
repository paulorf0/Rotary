package main

import (
	"encoding/json"
	"net/http"
	"time"

	"talos/db"
	"talos/providers"
)

// Onde vai ficar armazenado o pool ? O tempo de resposta para capturar um ip do pool e excluir um ip do pool não pode ser alto. Mas deixar em memória vai ser custoso se o sistema tiver muitos clientes
// type IpPool struct {
// }

type apiResponse struct {
	Msg    string `json:"msg"`
	Error  string `json:"error"`
	Status int    `json:"status"`
}

func main() {
	mux := http.NewServeMux()

	providers.AddClientPersistence = db.AddClient

	mux.HandleFunc("POST /v1/client/", v1ClientRoute)

	svr := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	svr.ListenAndServe()
}

func writeJSON[T any](w http.ResponseWriter, httpStatus int, value T) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	_ = json.NewEncoder(w).Encode(value)
}

func v1ClientRoute(w http.ResponseWriter, req *http.Request) {
	var v1ClientPayload struct {
		Client providers.Client `json:"client"`
	}

	if err := json.NewDecoder(req.Body).Decode(&v1ClientPayload); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "payload incorreto"})
		return
	}

	res := providers.AddClient(v1ClientPayload.Client)
	writeJSON(w, res.Status, res)
}
