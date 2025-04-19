package user

import (
	"context"

	"github.com/ogabrielrodrigues/imobiliary/internal/lib"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

func (s *Service) Authenticate(ctx context.Context, email, password string) (string, *response.Err) {
	found, err := s.repo.Authenticate(ctx, email, password)
	if err != nil {
		return "", err
	}

	token, err := lib.GenerateToken(found)
	if err != nil {
		return "", err
	}

	return token, nil
}
