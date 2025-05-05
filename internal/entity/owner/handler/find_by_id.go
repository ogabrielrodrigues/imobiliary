package owner

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/response"
)

func (h *Handler) FindByID(w http.ResponseWriter, r *http.Request) *response.Err {
	owner_id, err := uuid.Parse(r.PathValue("owner_id"))
	if err != nil {
		return response.NewErr(http.StatusBadRequest, response.ERR_INVALID_UUID)
	}

	owner, r_err := h.service.FindByID(r.Context(), owner_id)
	if r_err != nil {
		return r_err
	}

	return response.End(w, http.StatusOK, owner)
}
