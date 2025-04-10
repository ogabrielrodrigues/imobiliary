package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/environment"
	"github.com/ogabrielrodrigues/imobiliary/internal/api/dto"
	"github.com/ogabrielrodrigues/imobiliary/internal/api/entity"
	"github.com/ogabrielrodrigues/imobiliary/internal/shared/logger"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func Register(h *Handler, mux *http.ServeMux) {
	env := environment.LoadAPIEnvironment()
	// TODO: implementar lógica de gerenciamento de rotas
	users := []*entity.User{}

	id, _ := uuid.Parse("550e8400-e29b-41d4-a716-446655440000")
	user, _ := entity.NewUser(
		id,
		"67543-F",
		"João da Silva",
		"(11) 98765-4321",
		"joao.silva@imobdesk.com",
		"password",
		"https://github.com/ogabrielrodrigues.png",
	)
	users = append(users, user)

	mux.HandleFunc("POST /auth", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		logger.Info("auth", "request", r.Host)

		var dto dto.UserAuthDTO
		if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		for _, user := range users {
			if user.Email == dto.Email && user.ComparePwd(dto.Password) {
				token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
					"id":   user.ID,
					"user": user.ToDTO(),
					"exp":  time.Now().Add(time.Hour * 24 * 30).Unix(),
				})

				tokenString, err := token.SignedString([]byte(env.SECRET_KEY))
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				w.Header().Add("Authorization", fmt.Sprintf("Bearer %s", tokenString))
				w.WriteHeader(http.StatusOK)
				return
			}
		}

		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	})
}
