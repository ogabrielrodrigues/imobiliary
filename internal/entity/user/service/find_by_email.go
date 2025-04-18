package user

import (
	"context"

	"github.com/ogabrielrodrigues/imobiliary/internal/entity/user"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

func (s *Service) FindByEmail(ctx context.Context, email string) (*user.DTO, *response.Err) {
	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return user.ToDTO(), nil
}
