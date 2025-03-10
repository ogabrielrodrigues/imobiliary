package api

import (
	"net/http"

	"github.com/ogabrielrodrigues/imobiliary/internal/store/pg"
)

type handler struct {
	action *pg.Queries
	router *http.ServeMux
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

func Handler(queries *pg.Queries) http.Handler {
	h := handler{
		action: queries,
	}

	router := http.NewServeMux()

	h.router = router

	return h
}
