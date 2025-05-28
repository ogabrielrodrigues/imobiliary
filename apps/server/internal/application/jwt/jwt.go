package jwt

import (
	"net/http"
	"strings"
	"time"

	"imobiliary/config/environment"
	"imobiliary/internal/response"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

/*
	TODO: Check module logic
*/

func ExtractToken(authorization string) (string, error) {
	authorization = strings.TrimSpace(authorization)

	if !strings.HasPrefix(authorization, "Bearer") {
		return "", nil // TODO: place error type
	}

	token := strings.TrimPrefix(authorization, "Bearer")

	return token, nil
}

func GenerateToken(managerID uuid.UUID) (string, *response.Err) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": managerID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	token, err := claims.SignedString([]byte(environment.Environment.JWT_SECRET))
	if err != nil {
		return "", response.NewErr(http.StatusInternalServerError, err.Error()) // TODO: place error type
	}

	return token, nil
}

func ParseToken(token string) (uuid.UUID, *response.Err) {
	parsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(environment.Environment.JWT_SECRET), nil
	})

	if err != nil {
		return uuid.Nil, response.NewErr(http.StatusUnauthorized, err.Error()) // TODO: place error type
	}

	managerID, err := parsed.Claims.GetSubject()
	if err != nil {
		return uuid.Nil, response.NewErr(http.StatusUnauthorized, err.Error()) // TODO: place error type
	}

	return uuid.MustParse(managerID), nil
}
