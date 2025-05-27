package owner

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/owner"
	"github.com/ogabrielrodrigues/imobiliary/internal/response"
)

func (h *Handler) AssignOwnerToProperty(w http.ResponseWriter, r *http.Request) *response.Err {
	var dto owner.AssignOwnerDTO

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		return response.NewErr(http.StatusBadRequest, response.ERR_INVALID_REQUEST_BODY)
	}

	owner_id, err := uuid.Parse(dto.OwnerID)
	if err != nil {
		return response.NewErr(http.StatusBadRequest, response.ERR_INVALID_REQUEST_BODY)
	}

	property_id, err := uuid.Parse(dto.PropertyID)
	if err != nil {
		return response.NewErr(http.StatusBadRequest, response.ERR_INVALID_REQUEST_BODY)
	}

	if err := h.service.AssignOwnerToProperty(r.Context(), owner_id, property_id); err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	return nil
}
