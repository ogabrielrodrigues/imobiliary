package postgres

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Client struct {
	db *pgxpool.Pool
}

type Config struct {
	HealthCheckPeriod time.Duration
	MaxConnLifetime   time.Duration
	MaxConnIdleTime   time.Duration
}

func DefaultPostgresConfig() Config {
	return Config{
		HealthCheckPeriod: time.Minute,
		MaxConnLifetime:   30 * time.Minute,
		MaxConnIdleTime:   10 * time.Minute,
	}
}

func NewPostgresClient(connString string, config Config) (*pgxpool.Pool, error) {
	ctx := context.Background()

	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		return nil, err
	}

	pool.Config().HealthCheckPeriod = config.HealthCheckPeriod
	pool.Config().MaxConnLifetime = config.MaxConnLifetime
	pool.Config().MaxConnIdleTime = config.MaxConnIdleTime

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, err
	}

	return pool, nil
}
