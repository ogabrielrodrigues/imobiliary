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
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/property"
	property_handler "github.com/ogabrielrodrigues/imobiliary/internal/entity/property/handler"
	property_service "github.com/ogabrielrodrigues/imobiliary/internal/entity/property/service"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/user"
	"github.com/ogabrielrodrigues/imobiliary/internal/middleware"
	property_repository "github.com/ogabrielrodrigues/imobiliary/internal/provider/property/in_memory"
	"github.com/ogabrielrodrigues/imobiliary/internal/types"

	user_service "github.com/ogabrielrodrigues/imobiliary/internal/entity/user/service"
	avatar_repository "github.com/ogabrielrodrigues/imobiliary/internal/provider/avatar/in_memory"
	user_repository "github.com/ogabrielrodrigues/imobiliary/internal/provider/user/in_memory"
)

func TestE2ECreateProperty(t *testing.T) {
	environment.LoadFile(filepath.Join("..", "..", ".env"))

	pr := property_repository.NewInMemoryPropertyRepository()
	ps := property_service.NewService(pr)
	ph := property_handler.NewHandler(ps)

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

	t.Run("should be able to create a property", func(t *testing.T) {
		body := bytes.NewBuffer([]byte(`{
			"status": "Disponível",
			"water_id": "332345443@6",
			"energy_id": "234543566",
			"kind": "Residencial",
			"address": {
				"street": "Rua 07",
				"number": "448",
				"neighborhood": "Vl. Guarnieri",
				"complement": "",
				"city": "Colina",
				"state": "SP",
				"zip_code": "14770000"
			}
		}`))

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, "/properties", body)

		ctx := context.WithValue(request.Context(), middleware.UserIDKey, user_id.String())
		ph.Create(recorder, request.WithContext(ctx))

		response := recorder.Result()

		if response.StatusCode != http.StatusCreated {
			t.Errorf("expected status: %d, got: %d", http.StatusCreated, response.StatusCode)
		}
	})

	t.Run("should not be able to create a property with empty water id", func(t *testing.T) {
		body := bytes.NewBuffer([]byte(`{
			"status": "Disponível",
			"water_id": "",
			"energy_id": "234543566",
			"kind": "Residencial",
			"address": {
				"street": "Rua 07",
				"number": "448",
				"neighborhood": "Vl. Guarnieri",
				"complement": "",
				"city": "Colina",
				"state": "SP",
				"zip_code": "14770000"
			}
		}`))

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, "/properties", body)

		ctx := context.WithValue(request.Context(), middleware.UserIDKey, user_id.String())
		err := ph.Create(recorder, request.WithContext(ctx))

		if err.Code != http.StatusBadRequest {
			t.Errorf("expected status: %d, got: %d", http.StatusBadRequest, err.Code)
		}
	})

	t.Run("should not be able to create a property with empty energy id", func(t *testing.T) {
		body := bytes.NewBuffer([]byte(`{
			"status": "Disponível",
			"water_id": "332345443@6",
			"energy_id": "",
			"kind": "Residencial",
			"address": {
				"street": "Rua 07",
				"number": "448",
				"neighborhood": "Vl. Guarnieri",
				"complement": "",
				"city": "Colina",
				"state": "SP",
				"zip_code": "14770000"
			}
		}`))

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, "/properties", body)

		ctx := context.WithValue(request.Context(), middleware.UserIDKey, user_id.String())
		err := ph.Create(recorder, request.WithContext(ctx))

		if err.Code != http.StatusBadRequest {
			t.Errorf("expected status: %d, got: %d", http.StatusBadRequest, err.Code)
		}
	})

	t.Run("should not be able to create a property with empty address", func(t *testing.T) {
		body := bytes.NewBuffer([]byte(`{
			"status": "Disponível",
			"water_id": "332345443@6",
			"energy_id": "234543566",
			"kind": "Residencial"
		}`))

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, "/properties", body)

		ctx := context.WithValue(request.Context(), middleware.UserIDKey, user_id.String())
		err := ph.Create(recorder, request.WithContext(ctx))

		if err.Code != http.StatusBadRequest {
			t.Errorf("expected status: %d, got: %d", http.StatusBadRequest, err.Code)
		}
	})
}

func TestE2EFindByID(t *testing.T) {
	environment.LoadFile(filepath.Join("..", "..", ".env"))

	pr := property_repository.NewInMemoryPropertyRepository()
	ps := property_service.NewService(pr)
	ph := property_handler.NewHandler(ps)

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

	t.Run("should be able to find property by id", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "/properties/{property_id}", nil)
		request.SetPathValue("property_id", property_id.String())

		ctx := context.WithValue(request.Context(), middleware.UserIDKey, user_id.String())
		ph.FindByID(recorder, request.WithContext(ctx))

		response := recorder.Result()

		if response.StatusCode != http.StatusOK {
			t.Errorf("expected status: %d, got: %d", http.StatusCreated, response.StatusCode)
		}
	})

	t.Run("should not be able to find property with invalid id", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "/properties/{property_id}", nil)
		request.SetPathValue("property_id", "qualquercoisa")

		ctx := context.WithValue(request.Context(), middleware.UserIDKey, user_id.String())
		err := ph.FindByID(recorder, request.WithContext(ctx))

		if err.Code != http.StatusBadRequest {
			t.Errorf("expected status: %d, got: %d", http.StatusBadRequest, err.Code)
		}
	})

	t.Run("should not be able to find property if not exists", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "/properties/{property_id}", nil)
		request.SetPathValue("property_id", uuid.New().String())

		ctx := context.WithValue(request.Context(), middleware.UserIDKey, user_id.String())
		err := ph.FindByID(recorder, request.WithContext(ctx))

		if err.Code != http.StatusNotFound {
			t.Errorf("expected status: %d, got: %d", http.StatusNotFound, err.Code)
		}
	})

	t.Run("should not be able to find property without authentication", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "/properties/{property_id}", nil)
		request.SetPathValue("property_id", property_id.String())

		err := ph.FindByID(recorder, request)

		t.Log(err)

		if err.Code != http.StatusUnauthorized {
			t.Errorf("expected status: %d, got: %d", http.StatusUnauthorized, err.Code)
		}
	})
}

func TestE2EFindAllByUserID(t *testing.T) {
	environment.LoadFile(filepath.Join("..", "..", ".env"))

	pr := property_repository.NewInMemoryPropertyRepository()
	ps := property_service.NewService(pr)
	ph := property_handler.NewHandler(ps)

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

	t.Run("should be able to find all properties by user id", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "/properties", nil)

		ctx := context.WithValue(request.Context(), middleware.UserIDKey, user_id.String())
		ph.FindAllByUserID(recorder, request.WithContext(ctx))

		response := recorder.Result()

		if response.StatusCode != http.StatusOK {
			t.Errorf("expected status: %d, got: %d", http.StatusCreated, response.StatusCode)
		}
	})

	t.Run("should not be able to find all properties without authentication", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "/properties", nil)

		err := ph.FindAllByUserID(recorder, request)

		if err.Code != http.StatusUnauthorized {
			t.Errorf("expected status: %d, got: %d", http.StatusUnauthorized, err.Code)
		}
	})
}
