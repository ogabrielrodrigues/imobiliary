package main

import (
	"context"
	"errors"

	"net/http"
	"os"
	"os/signal"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ogabrielrodrigues/imobiliary/config/environment"
	"github.com/ogabrielrodrigues/imobiliary/config/logger"
	api "github.com/ogabrielrodrigues/imobiliary/internal"
	"github.com/ogabrielrodrigues/imobiliary/internal/middleware"
	"github.com/ogabrielrodrigues/imobiliary/internal/store"
)

func main() {
	env := environment.Load()

	pool, err := pgxpool.New(context.Background(), store.PGConnectionString(*env))
	if err != nil {
		logger.Log("error initializing database", err)
		return
	}

	handler := api.NewHandler(pool)

	go func() {
		logger.Log("server running on", env.SERVER_ADDR)
		if err := http.ListenAndServe(env.SERVER_ADDR, middleware.LoggerMiddleware(middleware.CORSMiddleware(handler))); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				logger.Log(err.Error())
				os.Exit(1)
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
