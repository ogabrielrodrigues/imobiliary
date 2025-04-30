package user

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/user"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

func (h *Handler) FindByID(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	id := r.PathValue("user_id")

	uid, u_err := uuid.Parse(id)
	if u_err != nil {
		err := response.NewErr(http.StatusBadRequest, user.ERR_UUID_INVALID)
		response.End(w, err.Code, err)
		return
	}

	user, err := h.service.FindByID(ctx, uid)
	if err != nil {
		response.End(w, err.Code, err)
		return
	}

	response.End(w, http.StatusOK, user)
}
