package plan

import (
	"context"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

type Service struct {
	repo IRepository
}

type IService interface {
	AssignPlanToUser(ctx context.Context, kind string, user_id uuid.UUID, plan *Plan) *response.Err
	GetUserPlan(ctx context.Context, user_id uuid.UUID) (*DTO, *response.Err)
}

func NewService(repo IRepository) *Service {
	return &Service{repo}
}

func (s *Service) AssignPlanToUser(ctx context.Context, kind string, user_id uuid.UUID, plan *Plan) *response.Err {
	return s.repo.AssignPlanToUser(ctx, kind, user_id, plan)
}

func (s *Service) GetUserPlan(ctx context.Context, user_id uuid.UUID) (*DTO, *response.Err) {
	plan, err := s.repo.GetUserPlan(ctx, user_id)
	if err != nil {
		return nil, err
	}

	return plan.ToDTO(), nil
}
