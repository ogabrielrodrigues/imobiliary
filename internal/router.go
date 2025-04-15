package api

import (
	"fmt"
	"net/http"

	"github.com/ogabrielrodrigues/imobiliary/internal/entity/user"
	repository "github.com/ogabrielrodrigues/imobiliary/internal/repository/user"
)

func Register(h *Handler, mux *http.ServeMux) {
	userHandler := user.NewHandler(
		user.NewService(
			repository.NewMemUserRepository(),
			repository.NewLocalUserAvatarRepository("./tmp"),
		),
	)

	mux.HandleFunc("GET /users", userHandler.FindBy)
	mux.HandleFunc("POST /users", userHandler.Create)
	mux.HandleFunc("PUT /users/{param}", userHandler.Update)
	mux.HandleFunc("DELETE /users/{id}", userHandler.Delete)
	mux.HandleFunc("POST /users/auth", userHandler.Authenticate)
	mux.HandleFunc("POST /users/avatar", userHandler.UpdateAvatar)

	mux.HandleFunc("GET /users/{user_id}/avatar", func(w http.ResponseWriter, r *http.Request) {
		user_id := r.PathValue("user_id")

		http.ServeFile(w, r, fmt.Sprintf("./tmp/%s/avatar.png", user_id))
	})
}
