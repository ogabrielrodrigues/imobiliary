package plan

import (
	"context"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/plan"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

func (s *Service) GetUserPlan(ctx context.Context, user_id uuid.UUID) (*plan.DTO, *response.Err) {
	plan, err := s.repo.GetUserPlan(ctx, user_id)
	if err != nil {
		return nil, err
	}

	return plan.ToDTO(), nil
}
