package api

import (
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ogabrielrodrigues/imobiliary/config/logger"
	"github.com/ogabrielrodrigues/imobiliary/internal/factory"
	"github.com/ogabrielrodrigues/imobiliary/internal/middleware"
)

func Register(h *Handler, mux *http.ServeMux, pool *pgxpool.Pool) {
	registerUserRoutes(mux, pool)

	registerPropertyRoutes(mux, pool)

	registerOwnerRoutes(mux, pool)
}

func registerUserRoutes(mux *http.ServeMux, pool *pgxpool.Pool) {
	user_handler, err := factory.NewUserHandlerFactory(pool)
	if err != nil {
		logger.Log(err)
		os.Exit(1)
	}

	mux.Handle("GET /users", http.HandlerFunc(user_handler.FindBy))
	mux.Handle("POST /users", http.HandlerFunc(user_handler.Create))
	mux.Handle("POST /users/auth", http.HandlerFunc(user_handler.Authenticate))

	mux.Handle("POST /users/avatar", middleware.AuthMiddleware(http.HandlerFunc(user_handler.UpdateAvatar)))
	mux.Handle("GET /users/plan", middleware.AuthMiddleware(http.HandlerFunc(user_handler.GetUserPlan)))
}

func registerPropertyRoutes(mux *http.ServeMux, pool *pgxpool.Pool) {
	property_handler := factory.NewPropertyHandlerFactory(pool)

	mux.Handle("GET /properties", middleware.AuthMiddleware(http.HandlerFunc(property_handler.FindAllByUserID)))
	mux.Handle("GET /properties/{property_id}", middleware.AuthMiddleware(http.HandlerFunc(property_handler.FindByID)))
	mux.Handle("POST /properties", middleware.AuthMiddleware(http.HandlerFunc(property_handler.Create)))
}

func registerOwnerRoutes(mux *http.ServeMux, pool *pgxpool.Pool) {
	owner_handler := factory.NewOwnerHandlerFactory(pool)

	mux.Handle("GET /owners/{owner_id}", middleware.AuthMiddleware(http.HandlerFunc(owner_handler.FindByID)))
	mux.Handle("GET /owners", middleware.AuthMiddleware(http.HandlerFunc(owner_handler.FindAllByManagerID)))
	mux.Handle("POST /owners", middleware.AuthMiddleware(http.HandlerFunc(owner_handler.Create)))
}
