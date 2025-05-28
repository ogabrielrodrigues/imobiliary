package owner

import (
	"context"

	"imobiliary/internal/entity/owner"
	"imobiliary/internal/response"

	"github.com/google/uuid"
)

func (s *Service) FindByID(ctx context.Context, owner_id uuid.UUID) (*owner.DTO, *response.Err) {
	return s.repo.FindByID(ctx, owner_id)
}
