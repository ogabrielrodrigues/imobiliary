package api

import (
	"net/http"

	"github.com/ogabrielrodrigues/imobiliary/internal/entity/user"
	repository "github.com/ogabrielrodrigues/imobiliary/internal/repository/user"
)

func Register(h *Handler, mux *http.ServeMux) {
	userHandler := user.NewHandler(
		user.NewService(
			repository.NewMemUserRepository(),
		),
	)

	mux.HandleFunc("GET /users", userHandler.FindBy)
	mux.HandleFunc("POST /users", userHandler.Create)
	mux.HandleFunc("PUT /users/{param}", userHandler.Update)
	mux.HandleFunc("DELETE /users/{id}", userHandler.Delete)
	mux.HandleFunc("POST /users/auth", userHandler.Authenticate)
}
