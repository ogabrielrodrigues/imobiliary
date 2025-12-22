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

type TenantHandler struct {
	findByIDTenantUseCase *usecase.FindByIDTenant
	createTenantUseCase   *usecase.CreateTenant
	findAllTenantUseCase  *usecase.FindAllTenant
}

func NewTenantHandler(
	findByIDTenantUseCase *usecase.FindByIDTenant,
	createTenantUseCase *usecase.CreateTenant,
	findAllTenantUseCase *usecase.FindAllTenant,
) *TenantHandler {
	return &TenantHandler{
		findByIDTenantUseCase,
		createTenantUseCase,
		findAllTenantUseCase,
	}
}

func (h *TenantHandler) FindByID(w http.ResponseWriter, r *http.Request) *httperr.HttpError {
	tenantIDParam := r.PathValue("tenant_id")
	tenantID, err := uuid.Parse(tenantIDParam)
	if err != nil {
		return httperr.NewBadRequestError(r.Context(), "invalid tenant id")
	}

	managerID, ok := r.Context().Value(middleware.ManagerIDKey).(string)
	if !ok {
		return httperr.NewUnauthorizedError(r.Context(), "not authorized")
	}

	tenantDTO, err := h.findByIDTenantUseCase.Execute(r.Context(), tenantID, uuid.MustParse(managerID))
	if err.(*httperr.HttpError) != nil {
		return err.(*httperr.HttpError)
	}

	return response.Json(w, http.StatusOK, tenantDTO)
}

func (h *TenantHandler) Create(w http.ResponseWriter, r *http.Request) *httperr.HttpError {
	var dto request.CreateTenantDTO

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		return httperr.NewBadRequestError(r.Context(), "error decoding json body")
	}

	managerID, ok := r.Context().Value(middleware.ManagerIDKey).(string)
	if !ok {
		return httperr.NewUnauthorizedError(r.Context(), "not authorized")
	}

	err := h.createTenantUseCase.Execute(r.Context(), dto, uuid.MustParse(managerID))
	if err != nil {
		return err
	}

	return response.End(w, http.StatusCreated)
}

func (h *TenantHandler) FindAll(w http.ResponseWriter, r *http.Request) *httperr.HttpError {
	managerID, ok := r.Context().Value(middleware.ManagerIDKey).(string)
	if !ok {
		return httperr.NewUnauthorizedError(r.Context(), "not authorized")
	}

	tenantsDTO, err := h.findAllTenantUseCase.Execute(r.Context(), uuid.MustParse(managerID))
	if err != nil {
		return err
	}

	return response.Json(w, http.StatusOK, tenantsDTO)
}
