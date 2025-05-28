package router

import (
	"imobiliary/config"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	pool   *pgxpool.Pool
	config *config.Config
	logger *logrus.Entry
	router *http.ServeMux
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

func NewRouter(pool *pgxpool.Pool, logger *logrus.Entry, config *config.Config) (http.Handler, error) {
	h := Handler{
		pool:   pool,
		config: config,
		logger: logger,
	}

	mux := http.NewServeMux()
	err := setupRoutes(&h)

	h.router = mux
	return h, err
}
