package owner

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ogabrielrodrigues/imobiliary/internal/entity/owner"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var dto owner.CreateDTO

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		r_err := response.NewErr(http.StatusBadRequest, owner.ERR_OWNER_BODY_INVALID)
		response.End(w, r_err.Code, r_err)
		return
	}

	id, err := h.service.Create(r.Context(), dto)
	if err != nil {
		response.End(w, err.Code, err)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("/owners/%s", id))
	w.WriteHeader(http.StatusCreated)
}
