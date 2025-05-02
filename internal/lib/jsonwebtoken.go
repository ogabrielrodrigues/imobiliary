package jwt

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/config/environment"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/user"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

const (
	ERR_TOKEN_INVALID_OR_EXPIRED string = "token inválido ou expirado"
)

func GenerateToken(user_id uuid.UUID) (string, *response.Err) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user_id,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	token, err := claims.SignedString([]byte(environment.Environment.SECRET_KEY))
	if err != nil {
		return "", response.NewErr(http.StatusInternalServerError, user.ERR_FAILED_GENERATE_TOKEN)
	}

	return token, nil
}

func ParseToken(token string) (uuid.UUID, *response.Err) {
	parsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(environment.Environment.SECRET_KEY), nil
	})

	if err != nil {
		return uuid.Nil, response.NewErr(http.StatusUnauthorized, ERR_TOKEN_INVALID_OR_EXPIRED)
	}

	user_id, err := parsed.Claims.GetSubject()
	if err != nil {
		return uuid.Nil, response.NewErr(http.StatusUnauthorized, ERR_TOKEN_INVALID_OR_EXPIRED)
	}

	return uuid.MustParse(user_id), nil
}
