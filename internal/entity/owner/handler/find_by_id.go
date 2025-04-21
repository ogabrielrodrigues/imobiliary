package owner

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

func (h *Handler) FindByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("owner_id")

	owner_id, err := uuid.Parse(id)
	if err != nil {
		r_err := response.NewErr(http.StatusBadRequest, err.Error())
		response.End(w, r_err.Code, r_err)
		return
	}

	owner, r_err := h.service.FindByID(r.Context(), owner_id)
	if r_err != nil {
		response.End(w, r_err.Code, r_err)
		return
	}

	response.End(w, http.StatusOK, owner)
}
