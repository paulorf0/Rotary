package main

import (
	"net/http"
	"time"
)

// Onde vai ficar armazenado o pool ? O tempo de resposta para capturar um ip do pool e excluir um ip do pool não pode ser alto. Mas deixar em memória vai ser custoso se o sistema tiver muitos clientes
// type IpPool struct {
// }

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, req *http.Request) {

	})

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
