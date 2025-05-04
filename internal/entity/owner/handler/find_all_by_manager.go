package owner

import (
	"net/http"

	"github.com/ogabrielrodrigues/imobiliary/internal/response"
)

func (h *Handler) FindAllByManagerID(w http.ResponseWriter, r *http.Request) {
	owners, err := h.service.FindAllByManagerID(r.Context())
	if err != nil {
		response.End(w, err.Code, err)
		return
	}

	response.End(w, http.StatusOK, owners)
}
