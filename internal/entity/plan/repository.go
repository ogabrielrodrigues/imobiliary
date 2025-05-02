package plan

import (
	"context"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

type IRepository interface {
	AssignPlanToUser(ctx context.Context, kind string, user_id uuid.UUID, plan *Plan) *response.Err
	GetUserPlan(ctx context.Context) (*Plan, *response.Err)
}
