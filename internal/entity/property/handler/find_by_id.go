package property

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/response"
)

func (h *Handler) FindByID(w http.ResponseWriter, r *http.Request) *response.Err {
	property_id := r.PathValue("property_id")

	uid, err := uuid.Parse(property_id)
	if err != nil {
		return response.NewErr(http.StatusBadRequest, err.Error())
	}

	property, r_err := h.service.FindByID(r.Context(), uid)
	if r_err != nil {
		return r_err
	}

	return response.End(w, http.StatusOK, property)
}
