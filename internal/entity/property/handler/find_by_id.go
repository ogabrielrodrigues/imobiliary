package property

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/response"
)

func (h *Handler) FindByID(w http.ResponseWriter, r *http.Request) {
	property_id := r.PathValue("property_id")

	uid, err := uuid.Parse(property_id)
	if err != nil {
		r_err := response.NewErr(http.StatusBadRequest, err.Error())
		response.End(w, r_err.Code, r_err)
		return
	}

	property, r_err := h.service.FindByID(r.Context(), uid)
	if r_err != nil {
		response.End(w, r_err.Code, r_err)
		return
	}

	response.End(w, http.StatusOK, property)
}
