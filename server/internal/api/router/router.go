package router

import (
	"imobiliary/config"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Handler struct {
	pool   *pgxpool.Pool
	config *config.Config
	logger *zap.Logger
	router *http.ServeMux
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

func NewRouter(pool *pgxpool.Pool, logger *zap.Logger, config *config.Config) (http.Handler, error) {
	h := Handler{
		pool:   pool,
		config: config,
		logger: logger,
	}

	mux := http.NewServeMux()
	err := setupRoutes(&h, mux)

	h.router = mux
	return h, err
}
