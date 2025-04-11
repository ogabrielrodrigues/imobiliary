package e2e

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	repository "github.com/ogabrielrodrigues/imobiliary/internal/api/repository/user"
	"github.com/ogabrielrodrigues/imobiliary/internal/api/user"
)

func TestCreateUser(t *testing.T) {
	userHandler := user.NewHandler(
		user.NewService(
			repository.NewMemUserRepository(),
		),
	)

	t.Run("should be able to create a user", func(t *testing.T) {
		body := bytes.NewBuffer([]byte(`{
			"creci_id": "12345-F",
			"fullname": "John Doe of Silva",
			"cellphone": "(11) 99999-9999"
			"email": "john.doe@example.com",
			"password": "password",
		}`))

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, "/users", body)

		userHandler.Create(recorder, request)

		response := recorder.Result()

		// TODO: unmarshal response body
		// body, _ := json.Unmarshal(response.Body, )
		// t.Log()

		if response.StatusCode != http.StatusCreated {
			t.Errorf("expected status: %d, got: %d", http.StatusCreated, response.StatusCode)
		}
	})
}
