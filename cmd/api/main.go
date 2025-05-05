package main

import (
	"context"
	"errors"
	"fmt"

	"net/http"
	"os"
	"os/signal"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ogabrielrodrigues/imobiliary/config/environment"
	"github.com/ogabrielrodrigues/imobiliary/config/logger"
	api "github.com/ogabrielrodrigues/imobiliary/internal"
	"github.com/ogabrielrodrigues/imobiliary/internal/middleware"
	store "github.com/ogabrielrodrigues/imobiliary/internal/store/postgres"
	"go.uber.org/zap"
)

func main() {
	env := environment.Load()

	pool, err := pgxpool.New(context.Background(), store.PostgresConnectionString(*env))
	if err != nil {
		logger.Error("error initializing database", zap.Error(err))
		return
	}

	handler := api.NewHandler(pool)

	go func() {
		logger.Info(fmt.Sprintf("server running on %s", env.SERVER_ADDR))
		if err := http.ListenAndServe(env.SERVER_ADDR, middleware.CORSMiddleware(middleware.LoggerMiddleware(handler))); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				logger.Error("error starting server", zap.Error(err))
				os.Exit(1)
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
