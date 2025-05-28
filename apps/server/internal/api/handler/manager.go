package handler

import (
	"encoding/json"
	"fmt"
	"imobiliary/internal/application/dto/request"
	"imobiliary/internal/application/usecase"
	"imobiliary/internal/response"
	"net/http"

	"github.com/google/uuid"
)

type ManagerHandler struct {
	findByIDManagerUseCase     *usecase.FindByIDManager
	createManagerUseCase       *usecase.CreateManager
	authenticateManagerUseCase *usecase.AuthenticateManager
}

func NewManagerHandler() *ManagerHandler {
	return &ManagerHandler{}
}

func (h *ManagerHandler) FindByID(w http.ResponseWriter, r *http.Request) *response.Err {
	managerID, u_err := uuid.Parse(r.PathValue("manager_id"))
	if u_err != nil {
		return response.NewErr(http.StatusBadRequest, response.ERR_INVALID_UUID)
	}

	managerDTO, err := h.findByIDManagerUseCase.Execute(r.Context(), managerID)
	if err != nil {
		return response.NewErr(http.StatusInternalServerError, err.Error()) // TODO: change error
	}

	return response.End(w, http.StatusOK, managerDTO)
}

func (h *ManagerHandler) Create(w http.ResponseWriter, r *http.Request) *response.Err {
	var dto request.CreateManagerDTO

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		return response.NewErr(http.StatusBadRequest, response.ERR_INVALID_REQUEST_BODY)
	}

	err := h.createManagerUseCase.Execute(r.Context(), dto)
	if err != nil {
		return response.NewErr(http.StatusInternalServerError, err.Error()) // TODO: change error
	}

	w.WriteHeader(http.StatusCreated)
	return nil
}

func (h *ManagerHandler) Authenticate(w http.ResponseWriter, r *http.Request) *response.Err {
	var dto request.AuthDTO

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		return response.NewErr(http.StatusBadRequest, response.ERR_INVALID_REQUEST_BODY)
	}

	token, err := h.authenticateManagerUseCase.Execute(r.Context(), dto)
	if err != nil {
		return response.NewErr(http.StatusInternalServerError, err.Error()) // TODO: change error
	}

	w.Header().Add("Authorization", fmt.Sprintf("Bearer %s", token))
	w.WriteHeader(http.StatusOK)
	return nil
}
