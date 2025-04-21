package provider

import (
	"context"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/plan"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

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
