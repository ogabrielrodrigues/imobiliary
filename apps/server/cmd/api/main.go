package main

import (
	"context"
	"fmt"
	"syscall"
	"time"

	"net/http"
	"os"
	"os/signal"

	"imobiliary/internal/api/middleware"
	"imobiliary/internal/api/router"
	"imobiliary/internal/application/logger"
	"imobiliary/internal/storage/postgres"

	"github.com/jackc/pgx/v5/pgxpool"
)

func initDependencies(config *Config) (*pgxpool.Pool, error) {
	postgresConnString := config.GetPostgresConnString()
	postgresConfig := postgres.DefaultPostgresConfig()

	postgresClient, err := postgres.NewPostgresClient(postgresConnString, postgresConfig)
	if err != nil {
		return nil, err
	}

	return postgresClient, nil
}

func main() {
	config, err := NewConfig()
	logger := logger.NewLogger(logger.Config{Environment: config.environment.GetEnvironment()})
	if err != nil {
		logger.Panic("config error", err)
	}

	ctx := context.Background()

	postgresClient, err := initDependencies(config)
	if err != nil {
		logger.Panic("error initializing dependencies", err)
	}

	handler, err := router.NewRouter(postgresClient)
	if err != nil {
		logger.Panic("error on routes setup", err)
	}

	server := &http.Server{
		Addr:           config.GetServerAddr(),
		Handler:        middleware.CORSMiddleware(middleware.LoggerMiddleware(handler)),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		logger.Info(fmt.Sprintf("server running on %s", config.GetServerAddr()))
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			logger.Panic("server error on listen", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	logger.Info("initialized graceful shutdown...")

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if server.Shutdown(ctx); err != nil {
		logger.Panic("fail to shutdown server", err)
	}

	logger.Info("server shutdown")
}
