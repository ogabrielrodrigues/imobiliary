package test

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/config/environment"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/owner"
	owner_handler "github.com/ogabrielrodrigues/imobiliary/internal/entity/owner/handler"
	owner_service "github.com/ogabrielrodrigues/imobiliary/internal/entity/owner/service"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/property"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/user"
	"github.com/ogabrielrodrigues/imobiliary/internal/middleware"
	owner_repository "github.com/ogabrielrodrigues/imobiliary/internal/provider/owner/in_memory"
	"github.com/ogabrielrodrigues/imobiliary/internal/response"
	"github.com/ogabrielrodrigues/imobiliary/internal/types"

	user_service "github.com/ogabrielrodrigues/imobiliary/internal/entity/user/service"
	avatar_repository "github.com/ogabrielrodrigues/imobiliary/internal/provider/avatar/in_memory"
	user_repository "github.com/ogabrielrodrigues/imobiliary/internal/provider/user/in_memory"

	property_repository "github.com/ogabrielrodrigues/imobiliary/internal/provider/property/in_memory"
)

func TestE2ECreateOwner(t *testing.T) {
	environment.LoadFile(filepath.Join("..", "..", ".env"))
	pr := property_repository.NewInMemoryPropertyRepository()
	or := owner_repository.NewInMemoryOwnerRepository(pr.GetProperties())
	os := owner_service.NewService(or)
	oh := owner_handler.NewHandler(os)

	t.Run("should be able to create a owner", func(t *testing.T) {
		body := bytes.NewBuffer([]byte(`{
			"fullname": "John Doe of Silva",
			"cpf": "123.456.768-90",
			"rg": "1234567-X",
			"email": "johndoe@example.com",
			"cellphone": "(00) 99000-0000",
			"occupation": "Developer",
			"marital_status": "Solteiro(a)",
			"address": {
				"street": "Rua da Saudade",
				"number": "099",
				"neighborhood": "Vl. Nobre",
				"complement": "APTO 07",
				"city": "São Paulo",
				"state": "SP",
				"zip_code": "14125070"
			}
		}`))

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, "/owners", body)

		oh.Create(recorder, request)

		response := recorder.Result()

		if response.StatusCode != http.StatusCreated {
			t.Errorf("expected status: %d, got: %d", http.StatusCreated, response.StatusCode)
		}
	})

	t.Run("should not be able to create a owner with empty name", func(t *testing.T) {
		body := bytes.NewBuffer([]byte(`{
			"fullname": "",
			"cpf": "123.456.768-90",
			"rg": "1234567-X",
			"email": "johndoe@example.com",
			"cellphone": "(00) 99000-0000",
			"occupation": "Developer",
			"marital_status": "Solteiro(a)",
			"address": {
				"street": "Rua da Saudade",
				"number": "099",
				"neighborhood": "Vl. Nobre",
				"complement": "APTO 07",
				"city": "São Paulo",
				"state": "SP",
				"zip_code": "14125070"
			}
		}`))

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, "/owners", body)

		err := oh.Create(recorder, request)

		if err.Message != owner.ERR_FULLNAME_EMPTY {
			t.Errorf("expected error: %s\ngot: %s", owner.ERR_FULLNAME_EMPTY, err.Message)
		}

		if err.Code != http.StatusBadRequest {
			t.Errorf("expected status: %d, got: %d", http.StatusBadRequest, err.Code)
		}
	})

	t.Run("should not be able to create a owner if name length less than 10 characters", func(t *testing.T) {
		body := bytes.NewBuffer([]byte(`{
			"fullname": "John",
			"cpf": "123.456.768-90",
			"rg": "1234567-X",
			"email": "johndoe@example.com",
			"cellphone": "(00) 99000-0000",
			"occupation": "Developer",
			"marital_status": "Solteiro(a)",
			"address": {
				"street": "Rua da Saudade",
				"number": "099",
				"neighborhood": "Vl. Nobre",
				"complement": "APTO 07",
				"city": "São Paulo",
				"state": "SP",
				"zip_code": "14125070"
			}
		}`))

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, "/owners", body)

		err := oh.Create(recorder, request)

		if err.Message != owner.ERR_FULLNAME_INVALID {
			t.Errorf("expected error: %s\ngot: %s", owner.ERR_FULLNAME_INVALID, err.Message)
		}

		if err.Code != http.StatusUnprocessableEntity {
			t.Errorf("expected status: %d, got: %d", http.StatusUnprocessableEntity, err.Code)
		}
	})

	t.Run("should not be able to create a owner if name length more than 100 characters", func(t *testing.T) {
		body := bytes.NewBuffer([]byte(`{
			"fullname": "John Doe of SilvaJohn Doe of SilvaJohn Doe of SilvaJohn Doe of SilvaJohn Doe of SilvaJohn Doe of SilvaJohn Doe of Silva",
			"cpf": "123.456.768-90",
			"rg": "1234567-X",
			"email": "johndoe@example.com",
			"cellphone": "(00) 99000-0000",
			"occupation": "Developer",
			"marital_status": "Solteiro(a)",
			"address": {
				"street": "Rua da Saudade",
				"number": "099",
				"neighborhood": "Vl. Nobre",
				"complement": "APTO 07",
				"city": "São Paulo",
				"state": "SP",
				"zip_code": "14125070"
			}
		}`))

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, "/owners", body)

		err := oh.Create(recorder, request)

		if err.Message != owner.ERR_FULLNAME_INVALID {
			t.Errorf("expected error: %s\ngot: %s", owner.ERR_FULLNAME_INVALID, err.Message)
		}

		if err.Code != http.StatusUnprocessableEntity {
			t.Errorf("expected status: %d, got: %d", http.StatusUnprocessableEntity, err.Code)
		}
	})

	t.Run("should not be able to create a owner with empty email", func(t *testing.T) {
		body := bytes.NewBuffer([]byte(`{
			"fullname": "John Doe of Silva",
			"cpf": "123.456.768-90",
			"rg": "1234567-X",
			"email": "",
			"cellphone": "(00) 99000-0000",
			"occupation": "Developer",
			"marital_status": "Solteiro(a)",
			"address": {
				"street": "Rua da Saudade",
				"number": "099",
				"neighborhood": "Vl. Nobre",
				"complement": "APTO 07",
				"city": "São Paulo",
				"state": "SP",
				"zip_code": "14125070"
			}
		}`))

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, "/owners", body)

		err := oh.Create(recorder, request)

		if err.Message != owner.ERR_EMAIL_EMPTY {
			t.Errorf("expected error: %s\ngot: %s", owner.ERR_EMAIL_EMPTY, err.Message)
		}

		if err.Code != http.StatusBadRequest {
			t.Errorf("expected status: %d, got: %d", http.StatusBadRequest, err.Code)
		}
	})

	t.Run("should not be able to create a owner if the email is invalid", func(t *testing.T) {
		body := bytes.NewBuffer([]byte(`{
			"fullname": "John Doe of Silva",
			"cpf": "123.456.768-90",
			"rg": "1234567-X",
			"email": "johndoeexample.com",
			"cellphone": "(00) 99000-0000",
			"occupation": "Developer",
			"marital_status": "Solteiro(a)",
			"address": {
				"street": "Rua da Saudade",
				"number": "099",
				"neighborhood": "Vl. Nobre",
				"complement": "APTO 07",
				"city": "São Paulo",
				"state": "SP",
				"zip_code": "14125070"
			}
		}`))

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, "/owners", body)

		err := oh.Create(recorder, request)

		if err.Message != owner.ERR_EMAIL_INVALID {
			t.Errorf("expected error: %s\ngot: %s", owner.ERR_EMAIL_INVALID, err.Message)
		}

		if err.Code != http.StatusUnprocessableEntity {
			t.Errorf("expected status: %d, got: %d", http.StatusUnprocessableEntity, err.Code)
		}
	})

	t.Run("should not be able to create a owner with empty cpf", func(t *testing.T) {
		body := bytes.NewBuffer([]byte(`{
			"fullname": "John Doe of Silva",
			"cpf": "",
			"rg": "1234567-X",
			"email": "johndoe@example.com",
			"cellphone": "(00) 99000-0000",
			"occupation": "Developer",
			"marital_status": "Solteiro(a)",
			"address": {
				"street": "Rua da Saudade",
				"number": "099",
				"neighborhood": "Vl. Nobre",
				"complement": "APTO 07",
				"city": "São Paulo",
				"state": "SP",
				"zip_code": "14125070"
			}
		}`))

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, "/owners", body)

		err := oh.Create(recorder, request)

		if err.Message != owner.ERR_CPF_EMPTY {
			t.Errorf("expected error: %s\ngot: %s", owner.ERR_CPF_EMPTY, err.Message)
		}

		if err.Code != http.StatusBadRequest {
			t.Errorf("expected status: %d, got: %d", http.StatusBadRequest, err.Code)
		}
	})

	t.Run("should not be able to create a owner if cpf is invalid", func(t *testing.T) {
		body := bytes.NewBuffer([]byte(`{
			"fullname": "John Doe of Silva",
			"cpf": "12345676890",
			"rg": "1234567-X",
			"email": "johndoe@example.com",
			"cellphone": "(00) 99000-0000",
			"occupation": "Developer",
			"marital_status": "Solteiro(a)",
			"address": {
				"street": "Rua da Saudade",
				"number": "099",
				"neighborhood": "Vl. Nobre",
				"complement": "APTO 07",
				"city": "São Paulo",
				"state": "SP",
				"zip_code": "14125070"
			}
		}`))

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, "/owners", body)

		err := oh.Create(recorder, request)

		if err.Message != owner.ERR_CPF_INVALID {
			t.Errorf("expected error: %s\ngot: %s", owner.ERR_CPF_INVALID, err.Message)
		}

		if err.Code != http.StatusUnprocessableEntity {
			t.Errorf("expected status: %d, got: %d", http.StatusUnprocessableEntity, err.Code)
		}
	})

	t.Run("should not be able to create a owner with empty rg", func(t *testing.T) {
		body := bytes.NewBuffer([]byte(`{
			"fullname": "John Doe of Silva",
			"cpf": "123.456.768-90",
			"rg": "",
			"email": "johndoe@example.com",
			"cellphone": "(00) 99000-0000",
			"occupation": "Developer",
			"marital_status": "Solteiro(a)",
			"address": {
				"street": "Rua da Saudade",
				"number": "099",
				"neighborhood": "Vl. Nobre",
				"complement": "APTO 07",
				"city": "São Paulo",
				"state": "SP",
				"zip_code": "14125070"
			}
		}`))

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, "/owners", body)

		err := oh.Create(recorder, request)

		if err.Message != owner.ERR_RG_EMPTY {
			t.Errorf("expected error: %s\ngot: %s", owner.ERR_RG_EMPTY, err.Message)
		}

		if err.Code != http.StatusBadRequest {
			t.Errorf("expected status: %d, got: %d", http.StatusBadRequest, err.Code)
		}
	})

	t.Run("should not be able to create a owner if rg length less than 5 characters", func(t *testing.T) {
		body := bytes.NewBuffer([]byte(`{
			"fullname": "John Doe of Silva",
			"cpf": "123.456.768-90",
			"rg": "1234",
			"email": "johndoe@example.com",
			"cellphone": "(00) 99000-0000",
			"occupation": "Developer",
			"marital_status": "Solteiro(a)",
			"address": {
				"street": "Rua da Saudade",
				"number": "099",
				"neighborhood": "Vl. Nobre",
				"complement": "APTO 07",
				"city": "São Paulo",
				"state": "SP",
				"zip_code": "14125070"
			}
		}`))

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, "/owners", body)

		err := oh.Create(recorder, request)

		if err.Message != owner.ERR_RG_INVALID {
			t.Errorf("expected error: %s\ngot: %s", owner.ERR_RG_INVALID, err.Message)
		}

		if err.Code != http.StatusUnprocessableEntity {
			t.Errorf("expected status: %d, got: %d", http.StatusUnprocessableEntity, err.Code)
		}
	})

	t.Run("should not be able to create a owner if rg length more than 15 characters", func(t *testing.T) {
		body := bytes.NewBuffer([]byte(`{
			"fullname": "John Doe of Silva",
			"cpf": "123.456.768-90",
			"rg": "1234567890123456789",
			"email": "johndoe@example.com",
			"cellphone": "(00) 99000-0000",
			"occupation": "Developer",
			"marital_status": "Solteiro(a)",
			"address": {
				"street": "Rua da Saudade",
				"number": "099",
				"neighborhood": "Vl. Nobre",
				"complement": "APTO 07",
				"city": "São Paulo",
				"state": "SP",
				"zip_code": "14125070"
			}
		}`))

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, "/owners", body)

		err := oh.Create(recorder, request)

		if err.Message != owner.ERR_RG_INVALID {
			t.Errorf("expected error: %s\ngot: %s", owner.ERR_RG_INVALID, err.Message)
		}

		if err.Code != http.StatusUnprocessableEntity {
			t.Errorf("expected status: %d, got: %d", http.StatusUnprocessableEntity, err.Code)
		}
	})

	t.Run("should not be able to create a owner with empty cellphone", func(t *testing.T) {
		body := bytes.NewBuffer([]byte(`{
			"fullname": "John Doe of Silva",
			"cpf": "123.456.768-90",
			"rg": "1234567-X",
			"email": "johndoe@example.com",
			"cellphone": "",
			"occupation": "Developer",
			"marital_status": "Solteiro(a)",
			"address": {
				"street": "Rua da Saudade",
				"number": "099",
				"neighborhood": "Vl. Nobre",
				"complement": "APTO 07",
				"city": "São Paulo",
				"state": "SP",
				"zip_code": "14125070"
			}
		}`))

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, "/owners", body)

		err := oh.Create(recorder, request)

		if err.Message != owner.ERR_CELLPHONE_EMPTY {
			t.Errorf("expected error: %s\ngot: %s", owner.ERR_CELLPHONE_EMPTY, err.Message)
		}

		if err.Code != http.StatusBadRequest {
			t.Errorf("expected status: %d, got: %d", http.StatusBadRequest, err.Code)
		}
	})

	t.Run("should not be able to create a owner if cellphone is invalid", func(t *testing.T) {
		body := bytes.NewBuffer([]byte(`{
			"fullname": "John Doe of Silva",
			"cpf": "123.456.768-90",
			"rg": "1234567-X",
			"email": "johndoe@example.com",
			"cellphone": "0099999999",
			"occupation": "Developer",
			"marital_status": "Solteiro(a)",
			"address": {
				"street": "Rua da Saudade",
				"number": "099",
				"neighborhood": "Vl. Nobre",
				"complement": "APTO 07",
				"city": "São Paulo",
				"state": "SP",
				"zip_code": "14125070"
			}
		}`))

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, "/owners", body)

		err := oh.Create(recorder, request)

		if err.Message != owner.ERR_CELLPHONE_INVALID {
			t.Errorf("expected error: %s\ngot: %s", owner.ERR_CELLPHONE_INVALID, err.Message)
		}

		if err.Code != http.StatusUnprocessableEntity {
			t.Errorf("expected status: %d, got: %d", http.StatusUnprocessableEntity, err.Code)
		}
	})
}

func TestE2EFindByIDUser(t *testing.T) {
	environment.LoadFile(filepath.Join("..", "..", ".env"))

	pr := property_repository.NewInMemoryPropertyRepository()
	or := owner_repository.NewInMemoryOwnerRepository(pr.GetProperties())
	os := owner_service.NewService(or)
	oh := owner_handler.NewHandler(os)

	owner_id, _ := os.Create(context.Background(), owner.CreateDTO{
		Fullname:      "John Doe of Silva",
		CPF:           "123.456.768-90",
		RG:            "1234567-X",
		Email:         "johndoe@example.com",
		Cellphone:     "(00) 99000-0000",
		Occupation:    "Developer",
		MaritalStatus: "Solteiro(a)",
		Address: types.AddressDTO{
			Street:       "Rua da Saudade",
			Number:       "099",
			Neighborhood: "Vl. Nobre",
			Complement:   "APTO 07",
			City:         "São Paulo",
			State:        "SP",
			ZipCode:      "14125070",
		},
	})

	t.Run("should not be able to find a owner with invalid id", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "/owners/{owner_id}", nil)
		request.SetPathValue("owner_id", "qualquercoisa")

		err := oh.FindByID(recorder, request)

		if err.Message != response.ERR_INVALID_UUID {
			t.Errorf("expected error: %s\ngot: %s", response.ERR_INVALID_UUID, err.Message)
		}
	})

	t.Run("should not be able to find a owner by id if not exists", func(t *testing.T) {
		id := uuid.New().String()

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "/owners/{owner_id}", nil)
		request.SetPathValue("owner_id", id)

		err := oh.FindByID(recorder, request)

		if err.Code != http.StatusNotFound {
			t.Errorf("expected status: %d, got: %d", http.StatusNotFound, err.Code)
		}
	})

	t.Run("should be able to find a user by id", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "/owners/{owner_id}", nil)
		request.SetPathValue("owner_id", owner_id.String())

		oh.FindByID(recorder, request)

		response := recorder.Result()

		if response.StatusCode != http.StatusOK {
			t.Errorf("expected status: %d, got: %d", http.StatusOK, response.StatusCode)
		}
	})
}

func TestE2EFindAllOwnersByManagerID(t *testing.T) {
	environment.LoadFile(filepath.Join("..", "..", ".env"))

	pr := property_repository.NewInMemoryPropertyRepository()
	or := owner_repository.NewInMemoryOwnerRepository(pr.GetProperties())
	os := owner_service.NewService(or)
	oh := owner_handler.NewHandler(os)

	ar := avatar_repository.NewInMemoryAvatarRepository("./tmp")
	ur := user_repository.NewInMemoryUserRepository()
	us := user_service.NewService(ur, ar)

	user_id, _ := us.Create(context.Background(), &user.CreateDTO{
		CreciID:   "12345-F",
		Fullname:  "John Doe of Silva",
		Cellphone: "(11) 99999-9999",
		Email:     "john.doe@example.com",
		Password:  "password",
	})

	or.Create(context.Background(), owner.Owner{
		ID:            uuid.New(),
		Fullname:      "John Doe of Silva",
		CPF:           "123.456.768-90",
		RG:            "1234567-X",
		Email:         "johndoe@example.com",
		Cellphone:     "(00) 99000-0000",
		Occupation:    "Developer",
		MaritalStatus: "Solteiro(a)",
		ManagerID:     user_id,
		Address: &types.Address{
			Street:       "Rua da Saudade",
			Number:       "099",
			Neighborhood: "Vl. Nobre",
			Complement:   "APTO 07",
			City:         "São Paulo",
			State:        "SP",
			ZipCode:      "14125070",
		},
	})

	t.Run("should be able to find all owners by manager id", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "/owners", nil)

		ctx := context.WithValue(request.Context(), middleware.UserIDKey, user_id.String())
		oh.FindAllByManagerID(recorder, request.WithContext(ctx))

		response := recorder.Result()

		if response.StatusCode != http.StatusOK {
			t.Errorf("expected status: %d, got: %d", http.StatusOK, response.StatusCode)
		}
	})

	t.Run("should not be able to find all owners without authentication", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "/owners", nil)

		err := oh.FindAllByManagerID(recorder, request)

		if err.Code != http.StatusUnauthorized {
			t.Errorf("expected status: %d, got: %d", http.StatusUnauthorized, err.Code)
		}
	})
}

func TestE2EAssignOwnerToProperty(t *testing.T) {
	environment.LoadFile(filepath.Join("..", "..", ".env"))

	pr := property_repository.NewInMemoryPropertyRepository()

	user_id := uuid.New()

	property_id := uuid.New()
	pr.Create(context.Background(), &property.Property{
		ID:       property_id,
		UserID:   user_id,
		Status:   property.StatusAvailable,
		WaterID:  "332345443@6",
		EnergyID: "234543566",
		Kind:     property.KindResidential,
		Address: &types.Address{
			Street:       "Rua 07",
			Number:       "448",
			Neighborhood: "Vl. Guarnieri",
			Complement:   "",
			City:         "Colina",
			State:        "SP",
			ZipCode:      "14770000",
		},
	})

	or := owner_repository.NewInMemoryOwnerRepository(pr.GetProperties())
	os := owner_service.NewService(or)
	oh := owner_handler.NewHandler(os)

	owner_id, _ := or.Create(context.Background(), owner.Owner{
		ID:            uuid.New(),
		Fullname:      "John Doe of Silva",
		CPF:           "123.456.768-90",
		RG:            "1234567-X",
		Email:         "johndoe@example.com",
		Cellphone:     "(00) 99000-0000",
		Occupation:    "Developer",
		MaritalStatus: "Solteiro(a)",
		ManagerID:     user_id,
		Address: &types.Address{
			Street:       "Rua da Saudade",
			Number:       "099",
			Neighborhood: "Vl. Nobre",
			Complement:   "APTO 07",
			City:         "São Paulo",
			State:        "SP",
			ZipCode:      "14125070",
		},
	})

	t.Run("should be able to assign owner to property", func(t *testing.T) {
		body := bytes.NewBuffer(fmt.Appendf(nil, `{
			"owner_id": "%s",
			"property_id": "%s"
		}`,
			owner_id.String(),
			property_id.String()))

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "/owners/assign", body)

		ctx := context.WithValue(request.Context(), middleware.UserIDKey, user_id.String())
		oh.AssignOwnerToProperty(recorder, request.WithContext(ctx))

		response := recorder.Result()

		if response.StatusCode != http.StatusOK {
			t.Errorf("expected status: %d, got: %d", http.StatusOK, response.StatusCode)
		}
	})

	t.Run("should not be able to assign owner to property with invalid owner_id", func(t *testing.T) {
		body := bytes.NewBuffer([]byte(fmt.Appendf(nil, `{
			"owner_id": "qualquercoisa",
			"property_id": "%s"
		}`,
			property_id.String())))

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "/owners/assign", body)

		ctx := context.WithValue(request.Context(), middleware.UserIDKey, user_id.String())
		err := oh.AssignOwnerToProperty(recorder, request.WithContext(ctx))

		fmt.Println(err.Message)

		if err.Code != http.StatusBadRequest {
			t.Errorf("expected status: %d, got: %d", http.StatusBadRequest, err.Code)
		}
	})

	t.Run("should not be able to assign owner to property with invalid property_id", func(t *testing.T) {
		body := bytes.NewBuffer(fmt.Appendf(nil, `{
			"owner_id": "%s",
			"property_id": "qualquercoisa"
		}`,
			owner_id.String()))

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "/owners/assign", body)

		ctx := context.WithValue(request.Context(), middleware.UserIDKey, user_id.String())
		err := oh.AssignOwnerToProperty(recorder, request.WithContext(ctx))

		if err.Code != http.StatusBadRequest {
			t.Errorf("expected status: %d, got: %d", http.StatusBadRequest, err.Code)
		}
	})

	t.Run("should not be able to assign owner to property with invalid owner_id", func(t *testing.T) {
		body := bytes.NewBuffer(fmt.Appendf(nil, `{
			"owner_id": "%s",
			"property_id": "%s"
		}`,
			uuid.New().String(),
			property_id.String()))

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "/owners/assign", body)

		ctx := context.WithValue(request.Context(), middleware.UserIDKey, user_id.String())
		err := oh.AssignOwnerToProperty(recorder, request.WithContext(ctx))

		if err.Code != http.StatusNotFound {
			t.Errorf("expected status: %d, got: %d", http.StatusNotFound, err.Code)
		}
	})

	t.Run("should not be able to assign owner to property with invalid property_id", func(t *testing.T) {
		body := bytes.NewBuffer(fmt.Appendf(nil, `{
			"owner_id": "%s",
			"property_id": "%s"
		}`,
			owner_id.String(),
			uuid.New().String()))

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "/owners/assign", body)

		ctx := context.WithValue(request.Context(), middleware.UserIDKey, user_id.String())
		err := oh.AssignOwnerToProperty(recorder, request.WithContext(ctx))

		if err.Code != http.StatusNotFound {
			t.Errorf("expected status: %d, got: %d", http.StatusNotFound, err.Code)
		}
	})

	t.Run("should not be able to assign owner to property without authorization", func(t *testing.T) {
		body := bytes.NewBuffer(fmt.Appendf(nil, `{
			"owner_id": "%s",
			"property_id": "%s"
		}`,
			owner_id.String(),
			property_id.String()))

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "/owners/assign", body)

		err := oh.AssignOwnerToProperty(recorder, request)

		if err.Code != http.StatusUnauthorized {
			t.Errorf("expected status: %d, got: %d", http.StatusUnauthorized, err.Code)
		}
	})
}
