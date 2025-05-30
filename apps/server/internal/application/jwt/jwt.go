package jwt

import (
	"context"
	"strings"
	"time"

	"imobiliary/internal/application/httperr"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

/*
	TODO: Check module logic
*/

func ExtractToken(authorization string) (string, *httperr.HttpError) {
	authorization = strings.TrimSpace(authorization)

	if !strings.HasPrefix(authorization, "Bearer") {
		return "", httperr.NewUnauthorizedError(context.Background(), "auth token invalid format")
	}

	token := strings.TrimPrefix(authorization, "Bearer")

	return token, nil
}

func GenerateToken(managerID uuid.UUID, jwtSecret string) (string, *httperr.HttpError) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": managerID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	token, err := claims.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", httperr.NewInternalServerError(context.Background(), "error generating jwt token")
	}

	return token, nil
}

func ParseToken(token string, jwtSecret string) (uuid.UUID, *httperr.HttpError) {
	parsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return uuid.Nil, httperr.NewUnauthorizedError(context.Background(), "error parsing jwt token")
	}

	managerID, err := parsed.Claims.GetSubject()
	if err != nil {
		return uuid.Nil, httperr.NewUnauthorizedError(context.Background(), "error retrieving jwt subject")
	}

	return uuid.MustParse(managerID), nil
}
