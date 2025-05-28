package owner

import (
	"net/http"

	"imobiliary/internal/response"
)

func (h *Handler) FindAllByManagerID(w http.ResponseWriter, r *http.Request) *response.Err {
	owners, err := h.service.FindAllByManagerID(r.Context())
	if err != nil {
		return err
	}

	if len(owners) == 0 {
		return response.End(w, http.StatusOK, []any{})
	}

	return response.End(w, http.StatusOK, owners)
}
