package plan

import (
	"context"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/plan"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

type Service struct {
	repo plan.IRepository
}

type IService interface {
	AssignPlanToUser(ctx context.Context, kind string, user_id uuid.UUID, plan *plan.Plan) *response.Err
	GetUserPlan(ctx context.Context, user_id uuid.UUID) (*plan.DTO, *response.Err)
}

func NewService(repo plan.IRepository) *Service {
	return &Service{repo}
}
