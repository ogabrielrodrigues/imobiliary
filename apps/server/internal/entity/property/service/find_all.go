package property

import (
	"context"

	"imobiliary/internal/entity/property"
	"imobiliary/internal/response"
)

func (s *Service) FindAllByUserID(ctx context.Context) ([]property.DTO, *response.Err) {
	return s.repo.FindAllByUserID(ctx)
}
