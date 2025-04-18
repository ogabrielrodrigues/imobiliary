package middleware

import (
	"net/http"

	"github.com/ogabrielrodrigues/imobiliary/config/environment"
)

func CORSMiddleware(next http.Handler) http.Handler {
	env := environment.Environment

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", env.CORS_ORIGIN)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		next.ServeHTTP(w, r)
	})
}
