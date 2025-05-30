package provider

import (
	"context"
	"net/http"

	"imobiliary/internal/response"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	ERR_CONNECT_TO_DB = "failed to connect to database"
)

type PostgresOwnerRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresOwnerRepository(pool *pgxpool.Pool) (*PostgresOwnerRepository, *response.Err) {
	if err := pool.Ping(context.Background()); err != nil {
		return nil, response.NewErr(http.StatusInternalServerError, ERR_CONNECT_TO_DB)
	}

	return &PostgresOwnerRepository{pool}, nil
}
