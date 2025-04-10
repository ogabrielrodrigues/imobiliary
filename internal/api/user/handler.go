package user

import (
	"context"
	"errors"
	"net/http"
	"regexp"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/api/response"
)

type Handler struct {
	service IService
}

func NewHandler(service IService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) FindBy(w http.ResponseWriter, r *http.Request) {
	param := r.PathValue("param")
	ctx := context.Background()

	id, err := uuid.Parse(param)
	if err != nil {
		if uuid.IsInvalidLengthError(err) {
			response.Error(w, http.StatusBadRequest, errors.New("invalid user id"))
			return
		}

		if match, _ := regexp.MatchString(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`, param); !match {
			response.Error(w, http.StatusBadRequest, errors.New("invalid user email"))
			return
		}

		user, err := h.service.FindByEmail(ctx, param)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, err)
			return
		}

		response.Json(w, http.StatusOK, user)
		return
	}

	h.service.FindByID(ctx, id)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	h.service.Create(nil, &User{})
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	h.service.Update(nil, &User{})
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	h.service.Delete(nil, &User{})
}
