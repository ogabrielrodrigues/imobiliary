package middleware

import (
	"fmt"
	"net/http"
	"time"

	"imobiliary/config/logger"

	"go.uber.org/zap"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		duration := uint(time.Since(start).Milliseconds())

		logger.Info(
			"request received",
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
			zap.String("duration", fmt.Sprintf("%dms", duration)),
		)
	})
}
