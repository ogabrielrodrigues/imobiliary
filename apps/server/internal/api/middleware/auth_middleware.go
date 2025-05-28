package middleware

import (
	"context"
	"net/http"

	"imobiliary/internal/application/jwt"
	"imobiliary/internal/response"
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
			err := response.NewErr(http.StatusUnauthorized, err.Error()) // TODO: place error type
			response.End(w, err.Code, err)
			return
		}

		managerID, err := jwt.ParseToken(token, jwtSecret)
		if err != nil {
			response.End(w, 401, err) // TODO: place error type
			return
		}

		ctx := context.WithValue(r.Context(), ManagerIDKey, managerID.String())

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
