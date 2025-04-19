package plan

import (
	"context"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/plan"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

func (s *Service) AssignPlanToUser(ctx context.Context, kind string, user_id uuid.UUID, plan *plan.Plan) *response.Err {
	return s.repo.AssignPlanToUser(ctx, kind, user_id, plan)
}
