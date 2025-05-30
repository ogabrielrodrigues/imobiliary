package property

import (
	"encoding/json"
	"net/http"

	"imobiliary/internal/entity/property"
	"imobiliary/internal/response"
)

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) *response.Err {
	var body *property.CreateDTO

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return response.NewErr(http.StatusBadRequest, response.ERR_INVALID_REQUEST_BODY)
	}

	if err := h.service.Create(r.Context(), body); err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	return nil
}
