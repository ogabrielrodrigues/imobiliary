package main

import (
	"context"
	"fmt"
	"syscall"
	"time"

	"net/http"
	"os"
	"os/signal"

	"imobiliary/config/logger"
	api "imobiliary/internal"
	"imobiliary/internal/middleware"
	"imobiliary/internal/storage/postgres"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
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
	if err != nil {
		logger.Panic("config error", zap.Error(err))
	}

	ctx := context.Background()

	postgresClient, err := initDependencies(config)
	if err != nil {
		logger.Panic("error initializing dependencies", zap.Error(err))
	}

	handler := api.NewHandler(postgresClient)

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
			logger.Panic("server error on listen", zap.Error(err))
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	logger.Info("initialized graceful shutdown...")

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if server.Shutdown(ctx); err != nil {
		logger.Panic("fail to shutdown server", zap.Error(err))
	}

	logger.Info("server shutdown")
}
