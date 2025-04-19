package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		duration := uint(time.Since(start).Milliseconds())

		fmt.Printf(
			"[%s] %dms > %s %s\n",
			time.Now().Format("01/02/2006 15:04:05"),
			duration,
			r.Method,
			r.URL.Path,
		)

		next.ServeHTTP(w, r)
	})
}
