package user

import (
	"net/http"

	"imobiliary/internal/entity/user"
	"imobiliary/internal/response"
)

func (h *Handler) ChangeAvatar(w http.ResponseWriter, r *http.Request) *response.Err {
	file, metadata, err := r.FormFile("avatar")
	if err != nil {
		return response.NewErr(http.StatusBadRequest, user.ERR_AVATAR_MUST_BE_PROVIDED)
	}

	if err := h.service.ChangeAvatar(r.Context(), file, metadata); err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	return nil
}
