package user

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/response"
)

func (h *Handler) FindByID(w http.ResponseWriter, r *http.Request) *response.Err {
	uid, u_err := uuid.Parse(r.PathValue("user_id"))
	if u_err != nil {
		return response.NewErr(http.StatusBadRequest, response.ERR_INVALID_UUID)
	}

	user, err := h.service.FindByID(r.Context(), uid)
	if err != nil {
		return err
	}

	return response.End(w, http.StatusOK, user)
}
