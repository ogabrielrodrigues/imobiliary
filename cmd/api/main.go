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
	"github.com/ogabrielrodrigues/imobiliary/internal/store"
)

func main() {
	env := environment.Load()

	pool, err := pgxpool.New(context.Background(), store.PGConnectionString(*env))
	if err != nil {
		logger.Log("pgxpool err=", err)
		return
	}

	handler := api.NewHandler(pool)

	go func() {
		if err := http.ListenAndServe(env.SERVER_ADDR, handler); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				panic(err)
			}

			logger.Log("Starting server on port " + env.SERVER_ADDR)
			if err := http.ListenAndServe(env.SERVER_ADDR, handler); err != nil {
				if !errors.Is(err, http.ErrServerClosed) {
					logger.Log("err", err)
					os.Exit(1)
				}
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
