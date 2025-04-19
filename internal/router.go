package api

import (
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ogabrielrodrigues/imobiliary/internal/factory"
	"github.com/ogabrielrodrigues/imobiliary/internal/middleware"
)

func Register(h *Handler, mux *http.ServeMux, pool *pgxpool.Pool) {
	registerUserRoutes(mux, pool)

	registerPropertyRoutes(mux)
}

func registerUserRoutes(mux *http.ServeMux, pool *pgxpool.Pool) {
	user_handler, err := factory.NewUserHandlerFactory(pool)
	if err != nil {
		os.Exit(1)
	}

	mux.Handle("GET /users", http.HandlerFunc(user_handler.FindBy))
	mux.Handle("POST /users", http.HandlerFunc(user_handler.Create))
	mux.Handle("POST /users/auth", http.HandlerFunc(user_handler.Authenticate))

	mux.Handle("POST /users/avatar", middleware.AuthMiddleware(http.HandlerFunc(user_handler.UpdateAvatar)))
	mux.Handle("GET /users/plan", middleware.AuthMiddleware(http.HandlerFunc(user_handler.GetUserPlan)))
}

func registerPropertyRoutes(mux *http.ServeMux) {
	property_handler := factory.NewPropertyHandlerFactory()

	mux.Handle("GET /properties", middleware.AuthMiddleware(http.HandlerFunc(property_handler.FindAllByUserID)))
	mux.Handle("GET /properties/{property_id}", middleware.AuthMiddleware(http.HandlerFunc(property_handler.FindByID)))
	mux.Handle("POST /properties", middleware.AuthMiddleware(http.HandlerFunc(property_handler.Create)))
}
