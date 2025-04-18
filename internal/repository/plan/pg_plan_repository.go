package repository

import (
	"context"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/plan"
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

func (pg *PostgresPlanRepository) AssignPlanToUser(ctx context.Context, kind string, user_id uuid.UUID, plan *plan.Plan) *response.Err {
	var query = `
	INSERT INTO "user_plan" (user_id, plan_id, properties_total_quota, properties_used_quota, properties_remaining_quota)
	VALUES ($1, $2, $3, $4, $5)`

	plan_id, err := pg.getPlanByKind(kind)
	if err != nil {
		return err
	}

	_, dberr := pg.pool.Exec(
		ctx,
		query,
		user_id,
		plan_id,
		plan.PropertiesTotalQuota,
		plan.PropertiesUsedQuota,
		plan.PropertiesRemainingQuota,
	)

	if dberr != nil {
		return response.NewErr(http.StatusInternalServerError, dberr.Error())
	}

	return nil
}

func (pg *PostgresPlanRepository) GetUserPlan(ctx context.Context, user_id uuid.UUID) (*plan.Plan, *response.Err) {
	tx, err := pg.pool.Begin(ctx)
	if err != nil {
		tx.Rollback(ctx)
		return nil, response.NewErr(http.StatusInternalServerError, err.Error())
	}

	var query = `
	SELECT p.id, p.kind, up.properties_total_quota, up.properties_used_quota, up.properties_remaining_quota
	FROM "user_plan" up
	JOIN "plan" p ON up.plan_id = p.id
	WHERE up.user_id = $1`

	row := tx.QueryRow(ctx, query, user_id)

	var plan plan.Plan
	if err := row.Scan(
		&plan.ID,
		&plan.Kind,
		&plan.PropertiesTotalQuota,
		&plan.PropertiesUsedQuota,
		&plan.PropertiesRemainingQuota,
	); err != nil {
		tx.Rollback(ctx)

		if errors.Is(err, pgx.ErrNoRows) {
			return nil, response.NewErr(http.StatusNotFound, ERR_PLAN_NOT_FOUND)
		}

		return nil, response.NewErr(http.StatusInternalServerError, err.Error())
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, response.NewErr(http.StatusInternalServerError, err.Error())
	}

	return &plan, nil
}
