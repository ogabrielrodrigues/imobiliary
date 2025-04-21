package owner

import (
	"context"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

func (s *Service) AssignOwnerToProperty(ctx context.Context, owner_id uuid.UUID, property_id uuid.UUID) *response.Err {
	return s.repo.AssignOwnerToProperty(ctx, owner_id, property_id)
}
