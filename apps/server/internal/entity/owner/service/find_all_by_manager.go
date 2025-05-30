package owner

import (
	"context"

	"imobiliary/internal/entity/owner"
	"imobiliary/internal/response"
)

func (s *Service) FindAllByManagerID(ctx context.Context) ([]owner.DTO, *response.Err) {
	return s.repo.FindAllByManagerID(ctx)
}
