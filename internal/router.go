package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ogabrielrodrigues/imobiliary/config/environment"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/user"
	repository "github.com/ogabrielrodrigues/imobiliary/internal/repository/user"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
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
	mux.HandleFunc("POST /users/avatar", userHandler.UpdateAvatar)

	mux.HandleFunc("GET /users/avatar/{avatar}", func(w http.ResponseWriter, r *http.Request) {
		avatar := r.PathValue("avatar")

		authorization, _ := strings.CutPrefix(r.Header.Get("Authorization"), "Bearer ")

		// TODO: fix this block of code
		token, err := jwt.Parse(authorization, func(token *jwt.Token) (interface{}, error) {
			return []byte(environment.Load().SECRET_KEY), nil
		})

		if err != nil {
			err := response.NewErr(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
			response.End(w, err.Code, err)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			user_id := claims["user_id"].(string)

			if user_id == "" {
				err := response.NewErr(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
				response.End(w, err.Code, err)
				return
			}

			http.ServeFile(w, r, fmt.Sprintf("./tmp/%s/%s", user_id, avatar))
		}
	})
}
