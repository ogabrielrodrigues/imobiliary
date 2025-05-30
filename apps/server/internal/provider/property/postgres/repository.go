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

type PostgresPropertyRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresPropertyRepository(pool *pgxpool.Pool) (*PostgresPropertyRepository, *response.Err) {
	if err := pool.Ping(context.Background()); err != nil {
		return nil, response.NewErr(http.StatusInternalServerError, ERR_CONNECT_TO_DB)
	}

	return &PostgresPropertyRepository{pool}, nil
}
