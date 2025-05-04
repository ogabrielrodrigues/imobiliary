package property

import (
	"encoding/json"
	"net/http"

	"github.com/ogabrielrodrigues/imobiliary/internal/entity/property"
	"github.com/ogabrielrodrigues/imobiliary/internal/response"
)

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var body *property.CreateDTO

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		r_err := response.NewErr(http.StatusBadRequest, err.Error())
		response.End(w, r_err.Code, r_err)
		return
	}

	r_err := h.service.Create(r.Context(), body)
	if r_err != nil {
		response.End(w, r_err.Code, r_err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
