package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

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

func IsUniqueConstraint(err error) bool {
	var pg_err *pgconn.PgError
	if !errors.As(err, &pg_err) {
		return false
	}

	if pg_err.Code == "23505" {
		return true
	}

	return false
}

func IsErrNoRows(err error) bool {
	return errors.Is(err, pgx.ErrNoRows)
}
