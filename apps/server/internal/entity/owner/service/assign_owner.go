package owner

import (
	"context"

	"imobiliary/internal/response"

	"github.com/google/uuid"
)

func (s *Service) AssignOwnerToProperty(ctx context.Context, owner_id uuid.UUID, property_id uuid.UUID) *response.Err {
	return s.repo.AssignOwnerToProperty(ctx, owner_id, property_id)
}
