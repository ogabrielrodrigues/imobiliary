package main

import (
	"errors"
	"net/http"
	"os"
	"os/signal"

	"github.com/ogabrielrodrigues/imobiliary/environment"
	"github.com/ogabrielrodrigues/imobiliary/internal/api"
	"github.com/ogabrielrodrigues/imobiliary/internal/api/middleware"
	"github.com/ogabrielrodrigues/imobiliary/internal/shared/logger"
)

func main() {
	env := environment.LoadAPIEnvironment()

	// ctx := context.Background()

	// pool, err := pgxpool.New(ctx, shared.ConnStr(env))
	// if err != nil {
	// 	logger.Error(logger.ErrDatabaseConnection, "err", err)
	// 	os.Exit(1)
	// }

	// defer pool.Close()

	// if err := pool.Ping(ctx); err != nil {
	// 	logger.Error(logger.ErrDatabaseConnection, "err", err)
	// 	os.Exit(1)
	// }

	handler := api.NewHandler(nil)

	go func() {
		logger.Info("Starting server on port " + env.SERVER_ADDR)
		if err := http.ListenAndServe(env.SERVER_ADDR, middleware.CORSMiddleware(env, handler)); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				logger.Error(logger.ErrInternalServer, "err", err)
				os.Exit(1)
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
