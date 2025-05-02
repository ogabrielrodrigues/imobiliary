package middleware

import (
	"net/http"
	"time"

	"github.com/ogabrielrodrigues/imobiliary/config/logger"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		duration := uint(time.Since(start).Milliseconds())

		logger.Logf(
			"[%s] %s > %dms\n",
			r.Method,
			r.URL.Path,
			duration,
		)
	})
}
