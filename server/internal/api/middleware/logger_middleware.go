package middleware

import (
	"net/http"
	"time"

	"github.com/phuslu/log"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		duration := uint(time.Since(start).Milliseconds())

		log.Info().Msgf("request received. %s %s ...%dms",
			r.Method,
			r.URL.Path,
			duration,
		)
	})
}
