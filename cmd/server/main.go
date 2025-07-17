package main

import (
	"net/http"
	"strings"

	"github.com/ricardolindner/go-expert-multithreading/internal/infra/webserver/handlers"
)

func main() {
	http.HandleFunc("/cep/", func(w http.ResponseWriter, r *http.Request) {
		cep := strings.TrimPrefix(r.URL.Path, "/cep/")
		if cep == "" {
			http.Error(w, "CEP not informed", http.StatusBadRequest)
			return
		}

		handlers.GetCepHandler(w, r, cep)
	})

	http.ListenAndServe(":8000", nil)
}
