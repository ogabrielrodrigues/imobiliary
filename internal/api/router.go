package api

import (
	"net/http"

	"github.com/ogabrielrodrigues/imobiliary/internal/api/service"
)

func Register(h *Handler, mux *http.ServeMux) {
	mux.HandleFunc("GET /rent", func(w http.ResponseWriter, r *http.Request) {
		service.RentMaker()
	})
	// TODO: implementar lógica de gerenciamento de rotas
}
