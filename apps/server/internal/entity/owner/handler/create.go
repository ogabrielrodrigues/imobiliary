package owner

import (
	"encoding/json"
	"fmt"
	"net/http"

	"imobiliary/internal/entity/owner"
	"imobiliary/internal/response"
)

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) *response.Err {
	var dto owner.CreateDTO

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		return response.NewErr(http.StatusBadRequest, response.ERR_INVALID_REQUEST_BODY)
	}

	id, err := h.service.Create(r.Context(), dto)
	if err != nil {
		return err
	}

	w.Header().Set("Location", fmt.Sprintf("/owners/%s", id))
	w.WriteHeader(http.StatusCreated)
	return nil
}
