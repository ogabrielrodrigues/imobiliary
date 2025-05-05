package api

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ogabrielrodrigues/imobiliary/config/logger"
	"github.com/ogabrielrodrigues/imobiliary/internal/factory"
	"github.com/ogabrielrodrigues/imobiliary/internal/middleware"
	"github.com/ogabrielrodrigues/imobiliary/internal/response"
	"go.uber.org/zap"
)

type RouteHandler func(w http.ResponseWriter, r *http.Request) *response.Err

func makeHandler(handler RouteHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := handler(w, r); err != nil {
			logger.Error(err.Message, zap.String("method", r.Method), zap.String("path", r.URL.Path), zap.Int("response_status", err.Code), zap.String("message", err.Message))
			response.End(w, err.Code, err)
		}
	}
}

func Register(h *Handler, mux *http.ServeMux, pool *pgxpool.Pool) {
	user_handler := factory.NewUserHandlerFactory(pool)

	mux.Handle("GET /users/{user_id}", makeHandler(user_handler.FindByID))
	mux.Handle("POST /users", makeHandler(user_handler.Create))
	mux.Handle("POST /users/auth", makeHandler(user_handler.Authenticate))
	mux.Handle("PUT /users/avatar", middleware.AuthMiddleware(makeHandler(user_handler.ChangeAvatar)))

	property_handler := factory.NewPropertyHandlerFactory(pool)

	mux.Handle("GET /properties", middleware.AuthMiddleware(makeHandler(property_handler.FindAllByUserID)))
	mux.Handle("GET /properties/{property_id}", middleware.AuthMiddleware(makeHandler(property_handler.FindByID)))
	mux.Handle("POST /properties", middleware.AuthMiddleware(makeHandler(property_handler.Create)))

	owner_handler := factory.NewOwnerHandlerFactory(pool)

	mux.Handle("GET /owners/{owner_id}", middleware.AuthMiddleware(makeHandler(owner_handler.FindByID)))
	mux.Handle("GET /owners", middleware.AuthMiddleware(makeHandler(owner_handler.FindAllByManagerID)))
	mux.Handle("POST /owners", middleware.AuthMiddleware(makeHandler(owner_handler.Create)))
	mux.Handle("PUT /owners/assign", middleware.AuthMiddleware(makeHandler(owner_handler.AssignOwnerToProperty)))
}
