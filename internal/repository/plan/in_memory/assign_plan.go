package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/plan"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

func (r *InMemoryPlanRepository) AssignPlanToUser(ctx context.Context, kind string, user_id uuid.UUID, plan *plan.Plan) *response.Err {
	r.plans[user_id.String()] = plan

	return nil
}
