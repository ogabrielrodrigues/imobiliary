package middleware

import (
	"net/http"

	"github.com/ogabrielrodrigues/imobiliary/internal/kind"
)

func CORSMiddleware(env *kind.APIEnvironment, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", env.CORS_ORIGIN)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		next.ServeHTTP(w, r)
	})
}
