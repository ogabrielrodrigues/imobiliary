package lib

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ogabrielrodrigues/imobiliary/config/environment"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/user"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

func GenerateToken(claim *user.User) (string, *response.Err) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  claim.ID,
		"user": claim.ToDTO(),
		"exp":  time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	token, t_err := claims.SignedString([]byte(environment.Environment.SECRET_KEY))
	if t_err != nil {
		return "", response.NewErr(http.StatusInternalServerError, user.ERR_FAILED_GENERATE_TOKEN)
	}

	return token, nil
}

func ParseToken(token string) (string, *response.Err) {
	parsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(environment.Environment.SECRET_KEY), nil
	})
	if err != nil {
		return "", response.NewErr(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
	}

	user_id, err := parsed.Claims.GetSubject()
	if err != nil {
		return "", response.NewErr(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
	}

	return user_id, nil
}

// func GetUserFromContext(ctx context.Context) (string, bool) {
// 	user_id, ok := ctx.Value(UserIDKey).(string)
// 	return user_id, ok
// }
