package user

import (
	"context"

	"imobiliary/internal/entity/user"
	"imobiliary/internal/response"
)

func (s *Service) ListAll(ctx context.Context) ([]user.DTO, *response.Err) {
	return s.repo.ListAll(ctx)
}
