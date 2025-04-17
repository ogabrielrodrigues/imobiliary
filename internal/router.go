package api

import (
	"net/http"

	"github.com/ogabrielrodrigues/imobiliary/internal/entity/property"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/user"
	"github.com/ogabrielrodrigues/imobiliary/internal/middleware"
	property_repository "github.com/ogabrielrodrigues/imobiliary/internal/repository/property"
	user_repository "github.com/ogabrielrodrigues/imobiliary/internal/repository/user"
)

func Register(h *Handler, mux *http.ServeMux) {
	registerUserRoutes(mux)

	registerPropertyRoutes(mux)
}

func registerUserRoutes(mux *http.ServeMux) {
	userHandler := user.NewHandler(
		user.NewService(
			user_repository.NewMemUserRepository(),
			user_repository.NewLocalUserAvatarRepository("./tmp"),
		),
	)

	mux.HandleFunc("GET /users", userHandler.FindBy)
	mux.HandleFunc("POST /users", userHandler.Create)
	mux.HandleFunc("PUT /users/{param}", userHandler.Update)
	mux.HandleFunc("DELETE /users/{id}", userHandler.Delete)
	mux.HandleFunc("POST /users/auth", userHandler.Authenticate)
	mux.HandleFunc("POST /users/avatar", userHandler.UpdateAvatar)
	mux.HandleFunc("GET /users/{user_id}/avatar", userHandler.GetAvatar)
}

func registerPropertyRoutes(mux *http.ServeMux) {
	propertyHandler := property.NewHandler(
		property.NewService(
			property_repository.NewMemPropertyRepository(),
		),
	)

	mux.Handle("GET /properties", middleware.AuthMiddleware(http.HandlerFunc(propertyHandler.FindAllByUserID)))
	mux.Handle("GET /properties/{property_id}", middleware.AuthMiddleware(http.HandlerFunc(propertyHandler.FindByID)))
	mux.Handle("POST /properties", middleware.AuthMiddleware(http.HandlerFunc(propertyHandler.Create)))
	mux.HandleFunc("PUT /properties/{property_id}", propertyHandler.Update)
	mux.HandleFunc("DELETE /properties/{property_id}", propertyHandler.Delete)
}
