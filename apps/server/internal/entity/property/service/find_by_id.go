package property

import (
	"context"

	"imobiliary/internal/entity/property"
	"imobiliary/internal/response"

	"github.com/google/uuid"
)

func (s *Service) FindByID(ctx context.Context, user_id uuid.UUID) (*property.DTO, *response.Err) {
	property, err := s.repo.FindByID(ctx, user_id)

	return property, err
}
