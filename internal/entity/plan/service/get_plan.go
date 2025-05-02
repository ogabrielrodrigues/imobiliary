package plan

import (
	"context"

	"github.com/ogabrielrodrigues/imobiliary/internal/entity/plan"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

func (s *Service) GetUserPlan(ctx context.Context) (*plan.DTO, *response.Err) {
	plan, err := s.repo.GetUserPlan(ctx)
	if err != nil {
		return nil, err
	}

	return plan.ToDTO(), nil
}
