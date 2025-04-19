package property

import (
	"context"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/property"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

func (s *Service) FindByID(ctx context.Context, user_id uuid.UUID) (*property.DTO, *response.Err) {
	property, err := s.repo.FindByID(ctx, user_id)

	return property, err
}
