package provider

import (
	"context"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

const (
	ERR_CONNECT_TO_DB  = "failed to connect to database"
	ERR_PLAN_NOT_FOUND = "plan not found"
)

type PostgresPlanRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresPlanRepository(pool *pgxpool.Pool) (*PostgresPlanRepository, *response.Err) {
	if err := pool.Ping(context.Background()); err != nil {
		return nil, response.NewErr(http.StatusInternalServerError, ERR_CONNECT_TO_DB)
	}

	return &PostgresPlanRepository{pool}, nil
}

func (pg *PostgresPlanRepository) getPlanByKind(kind string) (uuid.UUID, *response.Err) {
	var query = `SELECT id FROM "plan" WHERE kind = $1`

	row := pg.pool.QueryRow(context.Background(), query, kind)

	var plan_id string
	if err := row.Scan(&plan_id); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return uuid.Nil, response.NewErr(http.StatusNotFound, ERR_PLAN_NOT_FOUND)
		}

		return uuid.Nil, response.NewErr(http.StatusInternalServerError, err.Error())
	}

	return uuid.MustParse(plan_id), nil
}
