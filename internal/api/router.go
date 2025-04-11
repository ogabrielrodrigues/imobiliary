package api

import (
	"net/http"

	repository "github.com/ogabrielrodrigues/imobiliary/internal/api/repository/user"
	"github.com/ogabrielrodrigues/imobiliary/internal/api/user"
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
