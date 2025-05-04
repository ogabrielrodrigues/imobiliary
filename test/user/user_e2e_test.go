package test

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/config/environment"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/user"
	user_handler "github.com/ogabrielrodrigues/imobiliary/internal/entity/user/handler"
	user_service "github.com/ogabrielrodrigues/imobiliary/internal/entity/user/service"
	avatar_repository "github.com/ogabrielrodrigues/imobiliary/internal/provider/avatar/in_memory"
	user_repository "github.com/ogabrielrodrigues/imobiliary/internal/provider/user/in_memory"
)

func TestE2ECreateUser(t *testing.T) {
	environment.LoadFile(filepath.Join("..", "..", ".env"))

	ar := avatar_repository.NewInMemoryAvatarRepository("./tmp")
	ur := user_repository.NewInMemoryUserRepository()
	us := user_service.NewService(ur, ar)
	uh := user_handler.NewHandler(us)

	t.Run("should be able to create a user", func(t *testing.T) {
		body := bytes.NewBuffer([]byte(`{
			"creci_id": "12345-F",
			"fullname": "John Doe of Silva",
			"cellphone": "(11) 99999-9999",
			"email": "john.doe@example.com",
			"password": "password"
		}`))

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, "/users", body)

		uh.Create(recorder, request)

		response := recorder.Result()

		if response.StatusCode != http.StatusCreated {
			t.Errorf("expected status: %d, got: %d", http.StatusCreated, response.StatusCode)
		}
	})

	t.Run("should not be able to create a user with empty name", func(t *testing.T) {
		body := bytes.NewBuffer([]byte(`{
			"creci_id": "12345-F",
			"fullname": "",
			"cellphone": "(11) 99999-9999",
			"email": "john.doe@example.com",
			"password": "password"
		}`))

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, "/users", body)

		err := uh.Create(recorder, request)

		if err.Message != user.ERR_FULLNAME_EMPTY {
			t.Errorf("expected error: %s\ngot: %s", user.ERR_FULLNAME_EMPTY, err.Message)
		}

		if err.Code != http.StatusBadRequest {
			t.Errorf("expected status: %d, got: %d", http.StatusBadRequest, err.Code)
		}
	})

	t.Run("should not be able to create a user if name length less than 10 characters", func(t *testing.T) {
		body := bytes.NewBuffer([]byte(`{
			"creci_id": "12345-F",
			"fullname": "John Doe",
			"cellphone": "(11) 99999-9999",
			"email": "john.doe@example.com",
			"password": "password"
		}`))

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, "/users", body)

		err := uh.Create(recorder, request)

		if err.Message != user.ERR_FULLNAME_INVALID {
			t.Errorf("expected error: %s\ngot: %s", user.ERR_FULLNAME_INVALID, err.Message)
		}

		if err.Code != http.StatusUnprocessableEntity {
			t.Errorf("expected status: %d, got: %d", http.StatusUnprocessableEntity, err.Code)
		}
	})

	t.Run("should not be able to create a user if name length more than 100 characters", func(t *testing.T) {
		body := bytes.NewBuffer([]byte(`{
			"creci_id": "12345-F",
			"fullname": "John Doe of SilvaJohn Doe of SilvaJohn Doe of SilvaJohn Doe of SilvaJohn Doe of SilvaJohn Doeof Silva",
			"cellphone": "(11) 99999-9999",
			"email": "john.doe@example.com",
			"password": "password"
		}`))

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, "/users", body)

		err := uh.Create(recorder, request)

		if err.Message != user.ERR_FULLNAME_INVALID {
			t.Errorf("expected error: %s\ngot: %s", user.ERR_FULLNAME_INVALID, err.Message)
		}

		if err.Code != http.StatusUnprocessableEntity {
			t.Errorf("expected status: %d, got: %d", http.StatusUnprocessableEntity, err.Code)
		}
	})

	t.Run("should not be able to create a user with empty email", func(t *testing.T) {
		body := bytes.NewBuffer([]byte(`{
			"creci_id": "12345-F",
			"fullname": "John Doe of Silva",
			"cellphone": "(11) 99999-9999",
			"email": "",
			"password": "password"
		}`))

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, "/users", body)

		err := uh.Create(recorder, request)

		if err.Message != user.ERR_EMAIL_EMPTY {
			t.Errorf("expected error: %s\ngot: %s", user.ERR_EMAIL_EMPTY, err.Message)
		}

		if err.Code != http.StatusBadRequest {
			t.Errorf("expected status: %d, got: %d", http.StatusBadRequest, err.Code)
		}
	})

	t.Run("should not be able to create a user if the email is invalid", func(t *testing.T) {
		body := bytes.NewBuffer([]byte(`{
			"creci_id": "12345-F",
			"fullname": "John Doe of Silva",
			"cellphone": "(11) 99999-9999",
			"email": "invalidemail",
			"password": "password"
		}`))

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, "/users", body)

		err := uh.Create(recorder, request)

		if err.Message != user.ERR_EMAIL_INVALID {
			t.Errorf("expected error: %s\ngot: %s", user.ERR_EMAIL_INVALID, err.Message)
		}

		if err.Code != http.StatusUnprocessableEntity {
			t.Errorf("expected status: %d, got: %d", http.StatusUnprocessableEntity, err.Code)
		}
	})

	t.Run("should not be able to create a user with empty creci_id", func(t *testing.T) {
		body := bytes.NewBuffer([]byte(`{
			"creci_id": "",
			"fullname": "John Doe of Silva",
			"cellphone": "(11) 99999-9999",
			"email": "john.doe@example.com",
			"password": "password"
		}`))

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, "/users", body)

		err := uh.Create(recorder, request)

		if err.Message != user.ERR_CRECIID_EMPTY {
			t.Errorf("expected error: %s\ngot: %s", user.ERR_CRECIID_EMPTY, err.Message)
		}

		if err.Code != http.StatusBadRequest {
			t.Errorf("expected status: %d, got: %d", http.StatusBadRequest, err.Code)
		}
	})

	t.Run("should not be able to create a user if creci_id is invalid", func(t *testing.T) {
		body := bytes.NewBuffer([]byte(`{
			"creci_id": "21212F",
			"fullname": "John Doe of Silva",
			"cellphone": "(11) 99999-9999",
			"email": "john.doe@example.com",
			"password": "password"
		}`))

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, "/users", body)

		err := uh.Create(recorder, request)

		if err.Message != user.ERR_CRECIID_INVALID {
			t.Errorf("expected error: %s\ngot: %s", user.ERR_CRECIID_INVALID, err.Message)
		}

		if err.Code != http.StatusUnprocessableEntity {
			t.Errorf("expected status: %d, got: %d", http.StatusUnprocessableEntity, err.Code)
		}
	})

	t.Run("should not be able to create a user with empty cellphone", func(t *testing.T) {
		body := bytes.NewBuffer([]byte(`{
			"creci_id": "21212-F",
			"fullname": "John Doe of Silva",
			"cellphone": "",
			"email": "john.doe@example.com",
			"password": "password"
		}`))

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, "/users", body)

		err := uh.Create(recorder, request)

		if err.Message != user.ERR_CELLPHONE_EMPTY {
			t.Errorf("expected error: %s\ngot: %s", user.ERR_CELLPHONE_EMPTY, err.Message)
		}

		if err.Code != http.StatusBadRequest {
			t.Errorf("expected status: %d, got: %d", http.StatusBadRequest, err.Code)
		}
	})

	t.Run("should not be able to create a user if cellphone is invalid", func(t *testing.T) {
		body := bytes.NewBuffer([]byte(`{
			"creci_id": "21212-F",
			"fullname": "John Doe of Silva",
			"cellphone": "99999-9999",
			"email": "john.doe@example.com",
			"password": "password"
		}`))

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, "/users", body)

		err := uh.Create(recorder, request)

		if err.Message != user.ERR_CELLPHONE_INVALID {
			t.Errorf("expected error: %s\ngot: %s", user.ERR_CELLPHONE_INVALID, err.Message)
		}

		if err.Code != http.StatusUnprocessableEntity {
			t.Errorf("expected status: %d, got: %d", http.StatusUnprocessableEntity, err.Code)
		}
	})

	t.Run("should not be able to create a user with empty password", func(t *testing.T) {
		body := bytes.NewBuffer([]byte(`{
			"creci_id": "21212-F",
			"fullname": "John Doe of Silva",
			"cellphone": "(99) 9999-9999",
			"email": "john.doe@example.com",
			"password": ""
		}`))

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, "/users", body)

		err := uh.Create(recorder, request)

		if err.Message != user.ERR_PASSWORD_EMPTY {
			t.Errorf("expected error: %s\ngot: %s", user.ERR_PASSWORD_EMPTY, err.Message)
		}

		if err.Code != http.StatusBadRequest {
			t.Errorf("expected status: %d, got: %d", http.StatusBadRequest, err.Code)
		}
	})

	t.Run("should not be able to create a user if password is invalid", func(t *testing.T) {
		body := bytes.NewBuffer([]byte(`{
			"creci_id": "21212-F",
			"fullname": "John Doe of Silva",
			"cellphone": "(99) 99999-9999",
			"email": "john.doe@example.com",
			"password": "pwd"
		}`))

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, "/users", body)

		err := uh.Create(recorder, request)

		if err.Message != user.ERR_PASSWORD_INVALID {
			t.Errorf("expected error: %s\ngot: %s", user.ERR_PASSWORD_INVALID, err.Message)
		}

		if err.Code != http.StatusUnprocessableEntity {
			t.Errorf("expected status: %d, got: %d", http.StatusUnprocessableEntity, err.Code)
		}
	})
}

func TestE2EFindByIDUser(t *testing.T) {
	environment.LoadFile(filepath.Join("..", "..", ".env"))

	ar := avatar_repository.NewInMemoryAvatarRepository("./tmp")
	ur := user_repository.NewInMemoryUserRepository()
	us := user_service.NewService(ur, ar)
	uh := user_handler.NewHandler(us)

	t.Run("should not be able to find a user with invalid id", func(t *testing.T) {
		us.Create(context.Background(), &user.CreateDTO{
			CreciID:   "12345-F",
			Fullname:  "John Doe of Silva",
			Cellphone: "(11) 99999-9999",
			Email:     "john.doe@example.com",
			Password:  "password",
		})

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "/users/{user_id}", nil)
		request.SetPathValue("user_id", "qualquercoisa")

		err := uh.FindByID(recorder, request)

		if err.Message != user.ERR_UUID_INVALID {
			t.Errorf("expected error: %s\ngot: %s", user.ERR_UUID_INVALID, err.Message)
		}
	})

	t.Run("should not be able to find a user by id if not exists", func(t *testing.T) {
		id := uuid.New().String()

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "/users/{user_id}", nil)
		request.SetPathValue("user_id", id)

		err := uh.FindByID(recorder, request)

		if err.Code != http.StatusNotFound {
			t.Errorf("expected status: %d, got: %d", http.StatusNotFound, err.Code)
		}
	})

	t.Run("should be able to find a user by id", func(t *testing.T) {
		id, _ := us.Create(context.Background(), &user.CreateDTO{
			CreciID:   "12345-F",
			Fullname:  "John Doe of Silva",
			Cellphone: "(11) 99999-9999",
			Email:     "john.doe@example.com",
			Password:  "password",
		})

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "/users/{user_id}", nil)
		request.SetPathValue("user_id", id.String())

		uh.FindByID(recorder, request)

		response := recorder.Result()

		if response.StatusCode != http.StatusOK {
			t.Errorf("expected status: %d, got: %d", http.StatusOK, response.StatusCode)
		}
	})
}

func TestE2EAuthenticateUser(t *testing.T) {
	environment.LoadFile(filepath.Join("..", "..", ".env"))

	ar := avatar_repository.NewInMemoryAvatarRepository("./tmp")
	ur := user_repository.NewInMemoryUserRepository()
	us := user_service.NewService(ur, ar)
	uh := user_handler.NewHandler(us)

	t.Run("should be able to authenticate a user", func(t *testing.T) {
		us.Create(context.Background(), &user.CreateDTO{
			CreciID:   "12345-F",
			Fullname:  "John Doe of Silva",
			Cellphone: "(11) 99999-9999",
			Email:     "john.doe@example.com",
			Password:  "password",
		})

		body := bytes.NewBuffer([]byte(`{
			"email": "john.doe@example.com",
			"password": "password"
		}`))

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "/users/auth", body)

		uh.Authenticate(recorder, request)

		response := recorder.Result()

		if response.Header.Get("Authorization") == "" {
			t.Errorf("expected header: %s, got: %s", "Authorization", response.Header.Get("Authorization"))
		}

		if response.StatusCode != http.StatusOK {
			t.Errorf("expected status: %d, got: %d", http.StatusOK, response.StatusCode)
		}
	})

	t.Run("should not be able to authenticate a user with invalid body", func(t *testing.T) {
		us.Create(context.Background(), &user.CreateDTO{
			CreciID:   "12345-F",
			Fullname:  "John Doe of Silva",
			Cellphone: "(11) 99999-9999",
			Email:     "john.doe@example.com",
			Password:  "password",
		})

		body := bytes.NewBuffer([]byte(`{
			"email": 921921921,
			"password": "password"
		}`))

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "/users/auth", body)

		err := uh.Authenticate(recorder, request)

		if err.Code != http.StatusBadRequest {
			t.Errorf("expected status: %d, got: %d", http.StatusBadRequest, err.Code)
		}
	})

	t.Run("should not be able to authenticate a user with invalid email or password", func(t *testing.T) {
		body := bytes.NewBuffer([]byte(`{
			"email": "john.doe2@example.com",
			"password": "passworddd"
		}`))

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "/users/auth", body)

		err := uh.Authenticate(recorder, request)

		if err.Code != http.StatusNotFound {
			t.Errorf("expected status: %d, got: %d", http.StatusNotFound, err.Code)
		}
	})

	t.Run("should not be able to authenticate a user with invalid email or password", func(t *testing.T) {
		us.Create(context.Background(), &user.CreateDTO{
			CreciID:   "12345-F",
			Fullname:  "John Doe of Silva",
			Cellphone: "(11) 99999-9999",
			Email:     "john.doe@example.com",
			Password:  "password",
		})

		body := bytes.NewBuffer([]byte(`{
			"email": "john.doe@example.com",
			"password": "passworddd"
		}`))

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "/users/auth", body)

		err := uh.Authenticate(recorder, request)

		if err.Code != http.StatusUnauthorized {
			t.Errorf("expected status: %d, got: %d", http.StatusUnauthorized, err.Code)
		}
	})
}
