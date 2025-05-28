package user

import (
	"context"

	"imobiliary/internal/entity/user"
	jwt "imobiliary/internal/lib"
	"imobiliary/internal/response"
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
