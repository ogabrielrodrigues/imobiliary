package middleware

import (
	"context"
	"net/http"

	"imobiliary/internal/application/dto/response"
	"imobiliary/internal/application/jwt"
)

type contextKey string

func (c contextKey) String() string {
	return string(c)
}

var (
	ManagerIDKey = contextKey("manager_id")
)

func AuthMiddleware(next http.Handler, jwtSecret string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")

		token, err := jwt.ExtractToken(authorization)
		if err != nil {
			response.Json(w, err.HttpCode, err)
			return
		}

		managerID, err := jwt.ParseToken(token, jwtSecret)
		if err != nil {
			response.Json(w, err.HttpCode, err)
			return
		}

		ctx := context.WithValue(r.Context(), ManagerIDKey, managerID.String())

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
