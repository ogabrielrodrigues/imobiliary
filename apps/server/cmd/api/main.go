package main

import (
	"context"
	"errors"
	"fmt"
	"syscall"
	"time"

	"net/http"
	"os"
	"os/signal"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ogabrielrodrigues/imobiliary/config/environment"
	"github.com/ogabrielrodrigues/imobiliary/config/logger"
	api "github.com/ogabrielrodrigues/imobiliary/internal"
	"github.com/ogabrielrodrigues/imobiliary/internal/middleware"
	store "github.com/ogabrielrodrigues/imobiliary/internal/store/postgres"
	"github.com/ogabrielrodrigues/imobiliary/internal/types"
	"go.uber.org/zap"
)

func initDependencies(ctx context.Context, env *types.Environment) (*pgxpool.Pool, error) {
	return pgxpool.New(ctx, store.PostgresConnectionString(*env))
}

func main() {
	env := environment.Load()
	ctx := context.Background()

	pool, err := initDependencies(ctx, env)
	if err != nil {
		logger.Panic("error initializing database", zap.Error(err))
	}

	err = pool.Ping(ctx)
	if err != nil {
		logger.Panic("error connecting in database", zap.Error(err))
	}

	handler := api.NewHandler(pool)

	server := &http.Server{
		Addr:           env.SERVER_ADDR,
		Handler:        middleware.CORSMiddleware(middleware.LoggerMiddleware(handler)),
		ReadTimeout:    10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		logger.Info(fmt.Sprintf("server running on %s", env.SERVER_ADDR))
		if err := server.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				logger.Panic("error starting server", zap.Error(err))
			}

			logger.Error("server error on listen", zap.Error(err))
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if server.Shutdown(ctx); err != nil {
		logger.Panic("fail to shutdown server", zap.Error(err))
	}

	logger.Info("server shutdown")
}
