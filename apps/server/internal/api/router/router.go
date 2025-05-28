package router

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	pool   *pgxpool.Pool
	logger *logrus.Entry
	router *http.ServeMux
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

func NewRouter(pool *pgxpool.Pool, logger *logrus.Entry) (http.Handler, error) {
	h := Handler{
		pool:   pool,
		logger: logger,
	}

	mux := http.NewServeMux()
	err := setupRoutes(&h)

	h.router = mux
	return h, err
}
