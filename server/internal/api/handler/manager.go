package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"imobiliary/config"
	"imobiliary/internal/api/middleware"
	"imobiliary/internal/application/dto/request"
	"imobiliary/internal/application/dto/response"
	"imobiliary/internal/application/httperr"
	"imobiliary/internal/application/jwt"
	"imobiliary/internal/domain/manager"
	"imobiliary/internal/domain/types"
	"imobiliary/internal/mail"
	"net/http"

	"github.com/google/uuid"
)

type ManagerHandler struct {
	repository manager.Repository
	validator  *manager.Validator
	mailer     *mail.MailManagerRepository
}

func NewManagerHandler(
	repository manager.Repository,
	validator *manager.Validator,
	mailer *mail.MailManagerRepository,
) *ManagerHandler {
	return &ManagerHandler{
		repository,
		validator,
		mailer,
	}
}

func (h *ManagerHandler) FindByID(w http.ResponseWriter, r *http.Request) *httperr.HttpError {
	rawID, ok := r.Context().Value(middleware.ManagerIDKey).(string)
	if !ok {
		return httperr.NewUnauthorizedError(r.Context(), "not authorized")
	}

	managerID, err := uuid.Parse(rawID)
	if err != nil {
		return httperr.NewBadRequestError(r.Context(), "invalid managerID")
	}

	found, err := h.repository.FindByID(r.Context(), managerID)
	if err != nil {
		return err.(*httperr.HttpError)
	}

	managerDto := response.ManagerDto{
		ID:       found.ID.String(),
		Fullname: found.Fullname,
		Email:    found.Email.Value(),
		Phone:    found.Phone.Value(),
		Active:   found.Active,
	}

	return response.Json(w, http.StatusOK, managerDto)
}

func (h *ManagerHandler) Create(w http.ResponseWriter, r *http.Request) *httperr.HttpError {
	var dto request.CreateManagerDto

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		return httperr.NewBadRequestError(r.Context(), "error decoding json body")
	}

	newManager, err := h.validator.Validate(dto)
	if err != nil {
		var validationErrs *httperr.ValidationErrors
		if errors.Is(err, validationErrs) {
			return httperr.NewValidationError(r.Context(), validationErrs)
		}
		return err.(*httperr.HttpError)
	}

	if err := h.repository.Create(r.Context(), newManager); err != nil {
		return err
	}

	return response.End(w, http.StatusCreated)
}

func (h *ManagerHandler) Authenticate(w http.ResponseWriter, r *http.Request) *httperr.HttpError {
	config := config.GetConfig()

	var dto request.AuthDto

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		return httperr.NewBadRequestError(r.Context(), "error decoding json body")
	}

	email, err := types.NewEmail(dto.Email)
	if err != nil {
		return httperr.NewUnprocessableEntityError(r.Context(), err.Error())
	}

	managerID, err := h.repository.Authenticate(r.Context(), email, dto.Password)
	if err.(*httperr.HttpError) != nil {
		return err.(*httperr.HttpError)
	}

	token, err := jwt.GenerateToken(managerID, config.JwtSecret)
	if err.(*httperr.HttpError) != nil {
		return err.(*httperr.HttpError)
	}

	w.Header().Add("Authorization", fmt.Sprintf("Bearer %s", token))
	return response.End(w, http.StatusOK)
}

func (h *ManagerHandler) ResetPassword(w http.ResponseWriter, r *http.Request) *httperr.HttpError {
	var dto request.ResetPasswordDto

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		return httperr.NewBadRequestError(r.Context(), "error decoding json body")
	}

	email, err := types.NewEmail(dto.Email)
	if err != nil {
		return httperr.NewUnprocessableEntityError(r.Context(), err.Error())
	}

	found, err := h.repository.FindByEmail(r.Context(), email)
	if err.(*httperr.HttpError) != nil {
		return err.(*httperr.HttpError)
	}

	if found.Email.Value() != email.Value() {
		return httperr.NewForbiddenError(r.Context(), "this email doesn't belong to this manager.")
	}

	if err := h.mailer.ResetPasswordMessage(*email); err != nil {
		return err
	}

	return response.End(w, http.StatusOK)
}
