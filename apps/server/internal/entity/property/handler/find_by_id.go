package property

import (
	"net/http"

	"imobiliary/internal/response"

	"github.com/google/uuid"
)

func (h *Handler) FindByID(w http.ResponseWriter, r *http.Request) *response.Err {
	uid, err := uuid.Parse(r.PathValue("property_id"))
	if err != nil {
		return response.NewErr(http.StatusBadRequest, response.ERR_INVALID_UUID)
	}

	property, r_err := h.service.FindByID(r.Context(), uid)
	if r_err != nil {
		return r_err
	}

	return response.End(w, http.StatusOK, property)
}
