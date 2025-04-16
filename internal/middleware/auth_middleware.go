package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ogabrielrodrigues/imobiliary/config/environment"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") == "" {
			err := response.NewErr(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
			response.End(w, err.Code, err)
			return
		}

		authorization, _ := strings.CutPrefix(r.Header.Get("Authorization"), "Bearer ")

		token, err := jwt.Parse(authorization, func(token *jwt.Token) (interface{}, error) {
			return []byte(environment.Environment.SECRET_KEY), nil
		})
		if err != nil {
			err := response.NewErr(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
			response.End(w, err.Code, err)
			return
		}

		user_id, err := token.Claims.GetSubject()
		if err != nil {
			err := response.NewErr(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
			response.End(w, err.Code, err)
			return
		}

		req := r.WithContext(context.WithValue(r.Context(), "user_id", user_id))

		next.ServeHTTP(w, req)
	})
}
