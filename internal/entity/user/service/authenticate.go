package user

import (
	"context"

	"github.com/ogabrielrodrigues/imobiliary/internal/entity/user"
	jwt "github.com/ogabrielrodrigues/imobiliary/internal/lib"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

func (s *Service) Authenticate(ctx context.Context, dto *user.AuthDTO) (string, *response.Err) {
	found, err := s.repo.Authenticate(ctx, dto)
	if err != nil {
		return "", err
	}

	token, err := jwt.GenerateToken(found)
	if err != nil {
		return "", err
	}

	return token, nil
}
