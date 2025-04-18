package api

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct {
	pool   *pgxpool.Pool
	router *http.ServeMux
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

func NewHandler(pool *pgxpool.Pool) http.Handler {
	h := Handler{
		pool: pool,
	}

	mux := http.NewServeMux()
	Register(&h, mux, pool)

	h.router = mux
	return h
}
