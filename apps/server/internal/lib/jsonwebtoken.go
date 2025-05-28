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

const (
	ERR_TOKEN_INVALID_OR_EXPIRED string = "token inválido ou expirado"
	ERR_FAILED_GENERATE_TOKEN    string = "erro ao gerar o token de autenticação"
)

func ExtractToken(authorization string) string {
	if authorization == "" {
		return ""
	}

	authorization = strings.TrimSpace(authorization)

	if !strings.HasPrefix(authorization, "Bearer") {
		return ""
	}

	chunk := strings.Fields(authorization)

	if len(chunk) < 2 || chunk[1] == "" {
		return ""
	}

	token := chunk[1]

	token = strings.TrimSpace(token)
	if token == "" {
		return ""
	}

	return token
}

func GenerateToken(user_id uuid.UUID) (string, *response.Err) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user_id,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	token, err := claims.SignedString([]byte(environment.Environment.JWT_SECRET))
	if err != nil {
		return "", response.NewErr(http.StatusInternalServerError, ERR_FAILED_GENERATE_TOKEN)
	}

	return token, nil
}

func ParseToken(token string) (uuid.UUID, *response.Err) {
	parsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(environment.Environment.JWT_SECRET), nil
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
