package user

import (
	"net/http"

	"imobiliary/internal/response"
)

func (h *Handler) ListAll(w http.ResponseWriter, r *http.Request) *response.Err {
	users, err := h.service.ListAll(r.Context())
	if err != nil {
		return err
	}

	if len(users) == 0 {
		return response.End(w, http.StatusOK, []any{})
	}

	return response.End(w, http.StatusOK, users)
}
