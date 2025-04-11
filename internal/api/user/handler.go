package user

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/api/response"
	"github.com/ogabrielrodrigues/imobiliary/internal/shared/logger"
)

type Handler struct {
	service IService
}

func NewHandler(service IService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) FindBy(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	ctx := context.Background()

	email := params.Get("email")
	id := params.Get("id")

	logger.Info("Params", "params", params)

	if email != "" && id != "" {
		response.Error(w, http.StatusBadRequest, errors.New("only one parameter must be provided"))
		return
	}

	if id == "" {
		if match, _ := regexp.MatchString(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`, email); !match {
			response.Error(w, http.StatusBadRequest, errors.New("invalid user email"))
			return
		}

		user, err := h.service.FindByEmail(ctx, email)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, err)
			return
		}

		response.Json(w, http.StatusOK, user)
		return
	}

	if email == "" {
		uid, err := uuid.Parse(id)
		if err != nil {
			response.Error(w, http.StatusBadRequest, err)
			return
		}

		user, err := h.service.FindByID(ctx, uid)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, err)
			return
		}

		response.Json(w, http.StatusOK, user)
		return
	}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var dto CreateDTO
	ctx := context.Background()

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	id, err := h.service.Create(ctx, &dto)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("/users/%s", id))
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	h.service.Update(nil, &UpdateDTO{})
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	h.service.Delete(nil, uuid.MustParse(r.PathValue("id")))
}
