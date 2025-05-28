package middleware

import (
	"context"
	"net/http"
	"strings"

	jwt "imobiliary/internal/lib"
	"imobiliary/internal/response"
)

type contextKey string

func (c contextKey) String() string {
	return string(c)
}

var (
	UserIDKey = contextKey("user_id")
)

func clearMultipart(r *http.Request) {
	content_type := r.Header.Get("Content-Type")
	if strings.Contains(content_type, "multipart/form-data") {
		r.ParseMultipartForm(3 * 1024 * 1024)

		r.MultipartForm.RemoveAll()
	}
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")

		token := jwt.ExtractToken(authorization)
		if token == "" {
			clearMultipart(r)

			err := response.NewErr(http.StatusUnauthorized, jwt.ERR_TOKEN_INVALID_OR_EXPIRED)
			response.End(w, err.Code, err)
			return
		}

		user_id, err := jwt.ParseToken(token)
		if err != nil {
			clearMultipart(r)

			response.End(w, err.Code, err)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, user_id.String())

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
