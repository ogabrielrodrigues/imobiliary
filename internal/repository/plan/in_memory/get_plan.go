package repository

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/plan"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

func (r *InMemoryPlanRepository) GetUserPlan(ctx context.Context, user_id uuid.UUID) (*plan.Plan, *response.Err) {
	plan, ok := r.plans[user_id.String()]

	if !ok {
		return nil, response.NewErr(http.StatusNotFound, "plan not found")
	}

	return plan, nil
}
