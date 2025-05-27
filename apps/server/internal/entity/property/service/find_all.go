package property

import (
	"context"

	"github.com/ogabrielrodrigues/imobiliary/internal/entity/property"
	"github.com/ogabrielrodrigues/imobiliary/internal/response"
)

func (s *Service) FindAllByUserID(ctx context.Context) ([]property.DTO, *response.Err) {
	return s.repo.FindAllByUserID(ctx)
}
