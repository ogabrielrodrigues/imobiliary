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

type OwnerHandler struct {
	findByIDOwnerUseCase *usecase.FindByIDOwner
	createOwnerUseCase   *usecase.CreateOwner
	findAllOwnerUseCase  *usecase.FindAllOwner
}

func NewOwnerHandler(
	findByIDOwnerUseCase *usecase.FindByIDOwner,
	createOwnerUseCase *usecase.CreateOwner,
	findAllOwnerUseCase *usecase.FindAllOwner,
) *OwnerHandler {
	return &OwnerHandler{
		findByIDOwnerUseCase,
		createOwnerUseCase,
		findAllOwnerUseCase,
	}
}

func (h *OwnerHandler) FindByID(w http.ResponseWriter, r *http.Request) *httperr.HttpError {
	ownerIDParam := r.PathValue("owner_id")
	ownerID, err := uuid.Parse(ownerIDParam)
	if err != nil {
		return httperr.NewBadRequestError(r.Context(), "invalid owner id")
	}

	managerID, ok := r.Context().Value(middleware.ManagerIDKey).(string)
	if !ok {
		return httperr.NewUnauthorizedError(r.Context(), "not authorized")
	}

	ownerDTO, err := h.findByIDOwnerUseCase.Execute(r.Context(), ownerID, uuid.MustParse(managerID))
	if err.(*httperr.HttpError) != nil {
		return err.(*httperr.HttpError)
	}

	return response.Json(w, http.StatusOK, ownerDTO)
}

func (h *OwnerHandler) Create(w http.ResponseWriter, r *http.Request) *httperr.HttpError {
	var dto request.CreateOwnerDTO

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		return httperr.NewBadRequestError(r.Context(), "error decoding json body")
	}

	managerID, ok := r.Context().Value(middleware.ManagerIDKey).(string)
	if !ok {
		return httperr.NewUnauthorizedError(r.Context(), "not authorized")
	}

	err := h.createOwnerUseCase.Execute(r.Context(), dto, uuid.MustParse(managerID))
	if err != nil {
		return err
	}

	return response.End(w, http.StatusCreated)
}

func (h *OwnerHandler) FindAll(w http.ResponseWriter, r *http.Request) *httperr.HttpError {
	managerID, ok := r.Context().Value(middleware.ManagerIDKey).(string)
	if !ok {
		return httperr.NewUnauthorizedError(r.Context(), "not authorized")
	}

	ownersDTO, err := h.findAllOwnerUseCase.Execute(r.Context(), uuid.MustParse(managerID))
	if err != nil {
		return err
	}

	return response.Json(w, http.StatusOK, ownersDTO)
}
