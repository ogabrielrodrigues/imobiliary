package user

import (
	"context"

	"github.com/ogabrielrodrigues/imobiliary/internal/entity/user"
	"github.com/ogabrielrodrigues/imobiliary/internal/response"
)

func (s *Service) ListAll(ctx context.Context) ([]user.DTO, *response.Err) {
	return s.repo.ListAll(ctx)
}
