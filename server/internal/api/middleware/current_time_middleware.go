package middleware

import (
	"imobiliary/internal/application/context"
	"net/http"
	"time"
)

func CurrentTimeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.NewContextWithCurrentTime(r.Context(), time.Now().UTC())
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
