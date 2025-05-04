package property

import (
	"net/http"

	"github.com/ogabrielrodrigues/imobiliary/internal/response"
)

func (h *Handler) FindAllByUserID(w http.ResponseWriter, r *http.Request) *response.Err {
	properties, err := h.service.FindAllByUserID(r.Context())
	if err != nil {
		return err
	}

	if len(properties) == 0 {
		return response.End(w, http.StatusOK, []any{})
	}

	return response.End(w, http.StatusOK, properties)
}
