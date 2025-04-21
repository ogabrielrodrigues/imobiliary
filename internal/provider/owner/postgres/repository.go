package provider

import (
	"context"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
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
