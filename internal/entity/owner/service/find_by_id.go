package owner

import (
	"context"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/owner"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

func (s *Service) FindByID(ctx context.Context, owner_id uuid.UUID) (*owner.DTO, *response.Err) {
	return s.repo.FindByID(ctx, owner_id)
}
