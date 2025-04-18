package user

import (
	"context"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ogabrielrodrigues/imobiliary/config/environment"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/user"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

func (s *Service) Authenticate(ctx context.Context, email, password string) (string, *response.Err) {
	found, err := s.repo.Authenticate(ctx, email, password)
	if err != nil {
		return "", err
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   found.ID,
		"sub":  found.ID,
		"user": found.ToDTO(),
		"exp":  time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	token, t_err := claims.SignedString([]byte(environment.Environment.SECRET_KEY))
	if t_err != nil {
		return "", response.NewErr(http.StatusInternalServerError, user.ERR_FAILED_GENERATE_TOKEN)
	}

	return token, nil
}
