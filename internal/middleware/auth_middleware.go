package middleware

import (
	"context"
	"net/http"
	"strings"

	jwt "github.com/ogabrielrodrigues/imobiliary/internal/lib"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

type contextKey string

func (c contextKey) String() string {
	return string(c)
}

var (
	UserIDKey = contextKey("user_id")
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") == "" {
			err := response.NewErr(http.StatusUnauthorized, jwt.ERR_TOKEN_INVALID_OR_EXPIRED)
			response.End(w, err.Code, err)
			return
		}

		token := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")

		user_id, err := jwt.ParseToken(token)
		if err != nil {
			response.End(w, err.Code, err)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, user_id)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
