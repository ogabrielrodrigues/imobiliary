package middleware

import (
	"net/http"
	"time"

	"github.com/ogabrielrodrigues/imobiliary/config/logger"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		duration := uint(time.Since(start).Milliseconds())

		logger.Logf(
			"%dms > %s %s\n",
			duration,
			r.Method,
			r.URL.Path,
		)

		next.ServeHTTP(w, r)
	})
}
