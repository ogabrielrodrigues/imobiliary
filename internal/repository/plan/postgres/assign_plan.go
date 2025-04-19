package repository

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/plan"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

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
