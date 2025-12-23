package main

import (
	"context"
	"imobiliary/config"
	"imobiliary/internal/api/middleware"
	"imobiliary/internal/api/router"
	"imobiliary/internal/storage/postgres"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/phuslu/log"
)

func initDependencies(connString string) (*pgxpool.Pool, error) {
	postgresConfig := postgres.DefaultPostgresConfig()

	postgresClient, err := postgres.NewPostgresClient(connString, postgresConfig)
	if err != nil {
		return nil, err
	}

	return postgresClient, nil
}

func main() {
	err := config.LoadConfig()
	if err != nil {
		log.Panic().Msgf("error loading config. err=%v", err)
	}

	config := config.GetConfig()

	ctx := context.Background()

	postgresClient, err := initDependencies(config.PostgresConnString)
	if err != nil {
		log.Panic().Msgf("error initializing dependencies. err=%v", err)
	}

	handler := router.NewRouter(postgresClient)
	if err != nil {
		log.Panic().Msgf("error on routes setup. err=%v", err)
	}

	handler = middleware.CORSMiddleware(
		middleware.LoggerMiddleware(
			middleware.CurrentTimeMiddleware(handler),
		),
		config.CorsOrigin,
	)

	server := &http.Server{
		Addr:           config.ServerAddr,
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		log.Info().Msgf("server running on %s", config.ServerAddr)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Panic().Msgf("server error on listen. err=%v", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	log.Info().Msg("initialized graceful shutdown...")

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if server.Shutdown(ctx); err != nil {
		log.Panic().Msgf("fail to shutdown server. err=%v", err)
	}

	log.Info().Msg("server shutdown")
}
