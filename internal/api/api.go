package api

import (
	"net/http"

	"github.com/ogabrielrodrigues/imobiliary/internal/store/pg"
)

type Handler struct {
	query  *pg.Queries
	router *http.ServeMux
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

func NewHandler(queries *pg.Queries) http.Handler {
	h := Handler{
		query: queries,
	}

	mux := http.NewServeMux()
	Register(&h, mux)

	h.router = mux
	return h
}
