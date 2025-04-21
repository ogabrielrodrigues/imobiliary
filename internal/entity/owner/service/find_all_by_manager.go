package owner

import (
	"context"

	"github.com/ogabrielrodrigues/imobiliary/internal/entity/owner"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

func (s *Service) FindAllByManagerID(ctx context.Context) ([]owner.DTO, *response.Err) {
	return s.repo.FindAllByManagerID(ctx)
}
