package api

import (
	"net/http"

	repository "github.com/ogabrielrodrigues/imobiliary/internal/api/repository/user"
	"github.com/ogabrielrodrigues/imobiliary/internal/api/user"
)

func Register(h *Handler, mux *http.ServeMux) {
	// env := environment.LoadAPIEnvironment()
	// // TODO: implementar lógica de gerenciamento de rotas
	// users := []*user.User{}

	// id, _ := uuid.Parse("550e8400-e29b-41d4-a716-446655440000")
	// user, _ := user.New(
	// 	id,
	// 	"67543-F",
	// 	"João da Silva",
	// 	"(11) 98765-4321",
	// 	"joao.silva@imobdesk.com",
	// 	"password",
	// 	"https://github.com/ogabrielrodrigues.png",
	// )
	// users = append(users, user)

	// mux.HandleFunc("POST /auth", func(w http.ResponseWriter, r *http.Request) {
	// 	enableCors(&w)

	// 	logger.Info("auth", "request", r.Host)

	// 	var dto dto.UserAuthDTO
	// 	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
	// 		http.Error(w, err.Error(), http.StatusBadRequest)
	// 		return
	// 	}

	// 	for _, user := range users {
	// 		if user.Email == dto.Email && user.ComparePwd(dto.Password) {
	// 			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 				"id":   user.ID,
	// 				"user": user.ToDTO(),
	// 				"exp":  time.Now().Add(time.Hour * 24 * 30).Unix(),
	// 			})

	// 			tokenString, err := token.SignedString([]byte(env.SECRET_KEY))
	// 			if err != nil {
	// 				http.Error(w, err.Error(), http.StatusInternalServerError)
	// 				return
	// 			}

	// 			w.Header().Add("Authorization", fmt.Sprintf("Bearer %s", tokenString))
	// 			w.WriteHeader(http.StatusOK)
	// 			return
	// 		}
	// 	}

	// 	http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	// })

	userHandler := user.NewHandler(user.NewService(repository.NewMemUserRepository()))
	mux.HandleFunc("GET /users", userHandler.FindBy)
	mux.HandleFunc("POST /users", userHandler.Create)
	mux.HandleFunc("PUT /users/{param}", userHandler.Update)
	mux.HandleFunc("DELETE /users/{id}", userHandler.Delete)
}
