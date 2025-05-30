package handler

import (
	"encoding/json"
	"fmt"
	"imobiliary/internal/application/dto/request"
	"imobiliary/internal/application/dto/response"
	"imobiliary/internal/application/httperr"
	"imobiliary/internal/application/usecase"
	"net/http"

	"github.com/google/uuid"
)

type ManagerHandler struct {
	findByIDManagerUseCase     *usecase.FindByIDManager
	createManagerUseCase       *usecase.CreateManager
	authenticateManagerUseCase *usecase.AuthenticateManager
}

func NewManagerHandler(
	findByIDManagerUseCase *usecase.FindByIDManager,
	createManagerUseCase *usecase.CreateManager,
	authenticateManagerUseCase *usecase.AuthenticateManager) *ManagerHandler {
	return &ManagerHandler{
		findByIDManagerUseCase,
		createManagerUseCase,
		authenticateManagerUseCase,
	}
}

func (h *ManagerHandler) FindByID(w http.ResponseWriter, r *http.Request) *httperr.HttpError {
	managerID, u_err := uuid.Parse(r.PathValue("manager_id"))
	if u_err != nil {
		return httperr.NewBadRequestError(r.Context(), "invalid manager id param")
	}

	managerDTO, err := h.findByIDManagerUseCase.Execute(r.Context(), managerID)
	if err != nil {
		return err
	}

	return response.Json(w, http.StatusOK, managerDTO)
}

func (h *ManagerHandler) Create(w http.ResponseWriter, r *http.Request) *httperr.HttpError {
	var dto request.CreateManagerDTO

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		return httperr.NewBadRequestError(r.Context(), "error decoding json body")
	}

	err := h.createManagerUseCase.Execute(r.Context(), dto)
	if err != nil {
		return err
	}

	return response.End(w, http.StatusCreated)
}

func (h *ManagerHandler) Authenticate(w http.ResponseWriter, r *http.Request) *httperr.HttpError {
	var dto request.AuthDTO

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		return httperr.NewBadRequestError(r.Context(), "error decoding json body")
	}

	token, err := h.authenticateManagerUseCase.Execute(r.Context(), dto)
	if err != nil {
		return err
	}

	w.Header().Add("Authorization", fmt.Sprintf("Bearer %s", token))
	return response.End(w, http.StatusOK)
}
