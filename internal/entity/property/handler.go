package property

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

type Handler struct {
	service IService
}

func NewHandler(service IService) *Handler {
	return &Handler{service}
}

func (h *Handler) FindAllByUserID(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) FindByID(w http.ResponseWriter, r *http.Request) {
	property_id := r.PathValue("property_id")

	uid, err := uuid.Parse(property_id)
	if err != nil {
		r_err := response.NewErr(http.StatusBadRequest, err.Error())
		response.End(w, r_err.Code, r_err)
		return
	}

	property, r_err := h.service.FindByID(context.Background(), uid)
	if r_err != nil {
		response.End(w, r_err.Code, r_err)
		return
	}

	response.End(w, http.StatusOK, property)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var body *CreateDTO
	user_id := r.Context().Value("user_id").(string)

	uid, err := uuid.Parse(user_id)
	if err != nil {
		r_err := response.NewErr(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		response.End(w, r_err.Code, r_err)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		r_err := response.NewErr(http.StatusBadRequest, err.Error())
		response.End(w, r_err.Code, r_err)
		return
	}

	property, r_err := h.service.Create(context.Background(), body, uid)
	if r_err != nil {
		response.End(w, r_err.Code, r_err)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("/properties/%s", property.ID))
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
