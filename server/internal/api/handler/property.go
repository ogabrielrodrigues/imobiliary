package handler

import (
	"encoding/json"
	"imobiliary/internal/api/middleware"
	"imobiliary/internal/application/dto/request"
	"imobiliary/internal/application/dto/response"
	"imobiliary/internal/application/httperr"
	"imobiliary/internal/application/usecase"
	"net/http"

	"github.com/google/uuid"
)

type PropertyHandler struct {
	findByIDPropertyUseCase *usecase.FindByIDProperty
	createPropertyUseCase   *usecase.CreateProperty
	findAllPropertyUseCase  *usecase.FindAllProperty
}

func NewPropertyHandler(
	findByIDPropertyUseCase *usecase.FindByIDProperty,
	createPropertyUseCase *usecase.CreateProperty,
	findAllPropertyUseCase *usecase.FindAllProperty,
) *PropertyHandler {
	return &PropertyHandler{
		findByIDPropertyUseCase,
		createPropertyUseCase,
		findAllPropertyUseCase,
	}
}

func (h *PropertyHandler) FindByID(w http.ResponseWriter, r *http.Request) *httperr.HttpError {
	propertyIDParam := r.PathValue("property_id")
	propertyID, err := uuid.Parse(propertyIDParam)
	if err != nil {
		return httperr.NewBadRequestError(r.Context(), "invalid property id")
	}

	managerID, ok := r.Context().Value(middleware.ManagerIDKey).(string)
	if !ok {
		return httperr.NewUnauthorizedError(r.Context(), "not authorized")
	}

	propertyDTO, err := h.findByIDPropertyUseCase.Execute(r.Context(), propertyID, uuid.MustParse(managerID))
	if err.(*httperr.HttpError) != nil {
		return err.(*httperr.HttpError)
	}

	return response.Json(w, http.StatusOK, propertyDTO)
}

func (h *PropertyHandler) Create(w http.ResponseWriter, r *http.Request) *httperr.HttpError {
	var dto request.CreatePropertyDTO

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		return httperr.NewBadRequestError(r.Context(), "error decoding json body")
	}

	managerID, ok := r.Context().Value(middleware.ManagerIDKey).(string)
	if !ok {
		return httperr.NewUnauthorizedError(r.Context(), "not authorized")
	}

	err := h.createPropertyUseCase.Execute(r.Context(), dto, uuid.MustParse(managerID))
	if err != nil {
		return err
	}

	return response.End(w, http.StatusCreated)
}

func (h *PropertyHandler) FindAll(w http.ResponseWriter, r *http.Request) *httperr.HttpError {
	managerID, ok := r.Context().Value(middleware.ManagerIDKey).(string)
	if !ok {
		return httperr.NewUnauthorizedError(r.Context(), "not authorized")
	}

	propertiesDTO, err := h.findAllPropertyUseCase.Execute(r.Context(), uuid.MustParse(managerID))
	if err != nil {
		return err
	}

	return response.Json(w, http.StatusOK, propertiesDTO)
}
