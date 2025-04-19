package property

import (
	"net/http"

	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

func (h *Handler) FindAllByUserID(w http.ResponseWriter, r *http.Request) {
	properties, err := h.service.FindAllByUserID(r.Context())
	if err != nil {
		response.End(w, err.Code, err)
		return
	}

	if len(properties) == 0 {
		response.End(w, http.StatusOK, []any{})
		return
	}

	response.End(w, http.StatusOK, properties)
}
