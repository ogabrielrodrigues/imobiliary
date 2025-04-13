package test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joho/godotenv"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/user"
	repository "github.com/ogabrielrodrigues/imobiliary/internal/repository/user"
	res "github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

func TestE2ECreateUser(t *testing.T) {
	repo := repository.NewMemUserRepository()
	service := user.NewService(repo)
	userHandler := user.NewHandler(service)

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

		userHandler.Create(recorder, request)

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

		userHandler.Create(recorder, request)

		response := recorder.Result()

		var res_err res.Err
		json.NewDecoder(response.Body).Decode(&res_err)

		if res_err.Message != user.ERR_FULLNAME_EMPTY {
			t.Errorf("expected error: %s\ngot: %s", user.ERR_FULLNAME_EMPTY, res_err.Message)
		}

		if response.StatusCode != http.StatusBadRequest {
			t.Errorf("expected status: %d, got: %d", http.StatusBadRequest, response.StatusCode)
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

		userHandler.Create(recorder, request)

		response := recorder.Result()

		var res_err res.Err
		json.NewDecoder(response.Body).Decode(&res_err)

		if res_err.Message != user.ERR_FULLNAME_INVALID {
			t.Errorf("expected error: %s\ngot: %s", user.ERR_FULLNAME_INVALID, res_err.Message)
		}

		if response.StatusCode != http.StatusUnprocessableEntity {
			t.Errorf("expected status: %d, got: %d", http.StatusUnprocessableEntity, response.StatusCode)
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

		userHandler.Create(recorder, request)

		response := recorder.Result()

		var res_err res.Err
		json.NewDecoder(response.Body).Decode(&res_err)

		if res_err.Message != user.ERR_FULLNAME_INVALID {
			t.Errorf("expected error: %s\ngot: %s", user.ERR_FULLNAME_INVALID, res_err.Message)
		}

		if response.StatusCode != http.StatusUnprocessableEntity {
			t.Errorf("expected status: %d, got: %d", http.StatusUnprocessableEntity, response.StatusCode)
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

		userHandler.Create(recorder, request)

		response := recorder.Result()

		var res_err res.Err
		json.NewDecoder(response.Body).Decode(&res_err)

		if res_err.Message != user.ERR_EMAIL_EMPTY {
			t.Errorf("expected error: %s\ngot: %s", user.ERR_EMAIL_EMPTY, res_err.Message)
		}

		if response.StatusCode != http.StatusBadRequest {
			t.Errorf("expected status: %d, got: %d", http.StatusBadRequest, response.StatusCode)
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

		userHandler.Create(recorder, request)

		response := recorder.Result()

		var res_err res.Err
		json.NewDecoder(response.Body).Decode(&res_err)

		if res_err.Message != user.ERR_EMAIL_INVALID {
			t.Errorf("expected error: %s\ngot: %s", user.ERR_EMAIL_INVALID, res_err.Message)
		}

		if response.StatusCode != http.StatusUnprocessableEntity {
			t.Errorf("expected status: %d, got: %d", http.StatusUnprocessableEntity, response.StatusCode)
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

		userHandler.Create(recorder, request)

		response := recorder.Result()

		var res_err res.Err
		json.NewDecoder(response.Body).Decode(&res_err)

		if res_err.Message != user.ERR_CRECIID_EMPTY {
			t.Errorf("expected error: %s\ngot: %s", user.ERR_CRECIID_EMPTY, res_err.Message)
		}

		if response.StatusCode != http.StatusBadRequest {
			t.Errorf("expected status: %d, got: %d", http.StatusBadRequest, response.StatusCode)
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

		userHandler.Create(recorder, request)

		response := recorder.Result()

		var res_err res.Err
		json.NewDecoder(response.Body).Decode(&res_err)

		if res_err.Message != user.ERR_CRECIID_INVALID {
			t.Errorf("expected error: %s\ngot: %s", user.ERR_CRECIID_INVALID, res_err.Message)
		}

		if response.StatusCode != http.StatusUnprocessableEntity {
			t.Errorf("expected status: %d, got: %d", http.StatusUnprocessableEntity, response.StatusCode)
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

		userHandler.Create(recorder, request)

		response := recorder.Result()

		var res_err res.Err
		json.NewDecoder(response.Body).Decode(&res_err)

		if res_err.Message != user.ERR_CELLPHONE_EMPTY {
			t.Errorf("expected error: %s\ngot: %s", user.ERR_CELLPHONE_EMPTY, res_err.Message)
		}

		if response.StatusCode != http.StatusBadRequest {
			t.Errorf("expected status: %d, got: %d", http.StatusBadRequest, response.StatusCode)
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

		userHandler.Create(recorder, request)

		response := recorder.Result()

		var res_err res.Err
		json.NewDecoder(response.Body).Decode(&res_err)

		if res_err.Message != user.ERR_CELLPHONE_INVALID {
			t.Errorf("expected error: %s\ngot: %s", user.ERR_CELLPHONE_INVALID, res_err.Message)
		}

		if response.StatusCode != http.StatusUnprocessableEntity {
			t.Errorf("expected status: %d, got: %d", http.StatusUnprocessableEntity, response.StatusCode)
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

		userHandler.Create(recorder, request)

		response := recorder.Result()

		var res_err res.Err
		json.NewDecoder(response.Body).Decode(&res_err)

		if res_err.Message != user.ERR_PASSWORD_EMPTY {
			t.Errorf("expected error: %s\ngot: %s", user.ERR_PASSWORD_EMPTY, res_err.Message)
		}

		if response.StatusCode != http.StatusBadRequest {
			t.Errorf("expected status: %d, got: %d", http.StatusBadRequest, response.StatusCode)
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

		userHandler.Create(recorder, request)

		response := recorder.Result()

		var res_err res.Err
		json.NewDecoder(response.Body).Decode(&res_err)

		if res_err.Message != user.ERR_PASSWORD_INVALID {
			t.Errorf("expected error: %s\ngot: %s", user.ERR_PASSWORD_INVALID, res_err.Message)
		}

		if response.StatusCode != http.StatusUnprocessableEntity {
			t.Errorf("expected status: %d, got: %d", http.StatusUnprocessableEntity, response.StatusCode)
		}
	})
}

func TestE2EFindByUser(t *testing.T) {
	repo := repository.NewMemUserRepository()
	service := user.NewService(repo)
	userHandler := user.NewHandler(service)

	t.Run("should not be able to find a user by providing both ID and email together", func(t *testing.T) {
		id, _ := service.Create(context.Background(), &user.CreateDTO{
			CreciID:   "12345-F",
			Fullname:  "John Doe of Silva",
			Cellphone: "(11) 99999-9999",
			Email:     "john.doe@example.com",
			Password:  "password",
		})

		url := fmt.Sprintf("/users?id=%s&email=%s", id, "john.doe@example.com")

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, url, nil)

		userHandler.FindBy(recorder, request)

		response := recorder.Result()

		var res_err res.Err
		json.NewDecoder(response.Body).Decode(&res_err)

		if res_err.Message != user.ERR_ONLY_ONE_MUST_PARAMETER_MUST_BE_PROVIDED {
			t.Errorf("expected error: %s\ngot: %s", user.ERR_ONLY_ONE_MUST_PARAMETER_MUST_BE_PROVIDED, res_err.Message)
		}

		if response.StatusCode != http.StatusBadRequest {
			t.Errorf("expected status: %d, got: %d", http.StatusBadRequest, response.StatusCode)
		}
	})

	t.Run("should be able to find a user by email", func(t *testing.T) {
		service.Create(context.Background(), &user.CreateDTO{
			CreciID:   "12345-F",
			Fullname:  "John Doe of Silva",
			Cellphone: "(11) 99999-9999",
			Email:     "john.doe@example.com",
			Password:  "password",
		})

		url := fmt.Sprintf("/users?email=%s", "john.doe@example.com")

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, url, nil)

		userHandler.FindBy(recorder, request)

		response := recorder.Result()

		if response.StatusCode != http.StatusOK {
			t.Errorf("expected status: %d, got: %d", http.StatusOK, response.StatusCode)
		}
	})

	t.Run("should not be able to find a user with invalid email", func(t *testing.T) {
		service.Create(context.Background(), &user.CreateDTO{
			CreciID:   "12345-F",
			Fullname:  "John Doe of Silva",
			Cellphone: "(11) 99999-9999",
			Email:     "john.doe@example.com",
			Password:  "password",
		})

		url := fmt.Sprintf("/users?email=%s", "johnjohn")

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, url, nil)

		userHandler.FindBy(recorder, request)

		response := recorder.Result()

		var res_err res.Err
		json.NewDecoder(response.Body).Decode(&res_err)

		if res_err.Message != user.ERR_EMAIL_INVALID {
			t.Errorf("expected error: %s\ngot: %s", user.ERR_EMAIL_INVALID, res_err.Message)
		}

		if response.StatusCode != http.StatusBadRequest {
			t.Errorf("expected status: %d, got: %d", http.StatusBadRequest, response.StatusCode)
		}
	})

	t.Run("should not be able to find a user by email if not exists", func(t *testing.T) {
		url := fmt.Sprintf("/users?email=%s", "john.doe2@example.com")

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, url, nil)

		userHandler.FindBy(recorder, request)

		response := recorder.Result()

		if response.StatusCode != http.StatusNotFound {
			t.Errorf("expected status: %d, got: %d", http.StatusNotFound, response.StatusCode)
		}
	})

	t.Run("should be able to find a user by id", func(t *testing.T) {
		id, _ := service.Create(context.Background(), &user.CreateDTO{
			CreciID:   "12345-F",
			Fullname:  "John Doe of Silva",
			Cellphone: "(11) 99999-9999",
			Email:     "john.doe@example.com",
			Password:  "password",
		})

		url := fmt.Sprintf("/users?id=%s", id)

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, url, nil)

		userHandler.FindBy(recorder, request)

		response := recorder.Result()

		if response.StatusCode != http.StatusOK {
			t.Errorf("expected status: %d, got: %d", http.StatusOK, response.StatusCode)
		}
	})

	t.Run("should not be able to find a user with invalid id", func(t *testing.T) {
		service.Create(context.Background(), &user.CreateDTO{
			CreciID:   "12345-F",
			Fullname:  "John Doe of Silva",
			Cellphone: "(11) 99999-9999",
			Email:     "john.doe@example.com",
			Password:  "password",
		})

		url := fmt.Sprintf("/users?id=%s", "blablabla")

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, url, nil)

		userHandler.FindBy(recorder, request)

		response := recorder.Result()

		var res_err res.Err
		json.NewDecoder(response.Body).Decode(&res_err)

		if response.StatusCode != http.StatusBadRequest {
			t.Errorf("expected status: %d, got: %d", http.StatusBadRequest, response.StatusCode)
		}
	})

	t.Run("should not be able to find a user by id if not exists", func(t *testing.T) {
		service.Create(context.Background(), &user.CreateDTO{
			CreciID:   "12345-F",
			Fullname:  "John Doe of Silva",
			Cellphone: "(11) 99999-9999",
			Email:     "john.doe@example.com",
			Password:  "password",
		})

		url := fmt.Sprintf("/users?id=%s", "00000000-0000-0000-0000-000000000000")

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, url, nil)

		userHandler.FindBy(recorder, request)

		response := recorder.Result()

		if response.StatusCode != http.StatusNotFound {
			t.Errorf("expected status: %d, got: %d", http.StatusNotFound, response.StatusCode)
		}
	})
}

func TestE2EAuthenticateUser(t *testing.T) {
	godotenv.Load("../../../.env")

	repo := repository.NewMemUserRepository()
	service := user.NewService(repo)
	userHandler := user.NewHandler(service)

	t.Run("should be able to authenticate a user", func(t *testing.T) {
		service.Create(context.Background(), &user.CreateDTO{
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

		userHandler.Authenticate(recorder, request)

		response := recorder.Result()

		if response.Header.Get("Authorization") == "" {
			t.Errorf("expected header: %s, got: %s", "Authorization", response.Header.Get("Authorization"))
		}

		if response.StatusCode != http.StatusOK {
			t.Errorf("expected status: %d, got: %d", http.StatusOK, response.StatusCode)
		}
	})

	t.Run("should not be able to authenticate a user with invalid body", func(t *testing.T) {
		service.Create(context.Background(), &user.CreateDTO{
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

		userHandler.Authenticate(recorder, request)

		response := recorder.Result()

		if response.StatusCode != http.StatusBadRequest {
			t.Errorf("expected status: %d, got: %d", http.StatusBadRequest, response.StatusCode)
		}
	})

	t.Run("should not be able to authenticate a user with invalid email or password", func(t *testing.T) {
		body := bytes.NewBuffer([]byte(`{
			"email": "john.doe2@example.com",
			"password": "passworddd"
		}`))

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "/users/auth", body)

		userHandler.Authenticate(recorder, request)

		response := recorder.Result()

		if response.StatusCode != http.StatusNotFound {
			t.Errorf("expected status: %d, got: %d", http.StatusNotFound, response.StatusCode)
		}
	})

	t.Run("should not be able to authenticate a user with invalid email or password", func(t *testing.T) {
		service.Create(context.Background(), &user.CreateDTO{
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

		userHandler.Authenticate(recorder, request)

		response := recorder.Result()

		if response.StatusCode != http.StatusUnauthorized {
			t.Errorf("expected status: %d, got: %d", http.StatusUnauthorized, response.StatusCode)
		}
	})
}
