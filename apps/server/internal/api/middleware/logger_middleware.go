package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

func LoggerMiddleware(next http.Handler, logger *logrus.Entry) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		duration := uint(time.Since(start).Milliseconds())

		logger.Info(
			"request received",
			logrus.WithFields(logrus.Fields{
				"method":   r.Method,
				"path":     r.URL.Path,
				"duration": fmt.Sprintf("%dms", duration),
			}),
		)
	})
}
