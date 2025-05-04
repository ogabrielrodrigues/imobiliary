package owner

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/owner"
	"github.com/ogabrielrodrigues/imobiliary/internal/response"
)

func (h *Handler) AssignOwnerToProperty(w http.ResponseWriter, r *http.Request) {
	var dto owner.AssignOwnerDTO

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		r_err := response.NewErr(http.StatusBadRequest, owner.ERR_OWNER_BODY_INVALID)
		response.End(w, r_err.Code, r_err)
		return
	}

	owner_id, err := uuid.Parse(dto.OwnerID)
	if err != nil {
		r_err := response.NewErr(http.StatusBadRequest, err.Error())
		response.End(w, r_err.Code, r_err)
		return
	}

	property_id, err := uuid.Parse(dto.PropertyID)
	if err != nil {
		r_err := response.NewErr(http.StatusBadRequest, err.Error())
		response.End(w, r_err.Code, r_err)
		return
	}

	if err := h.service.AssignOwnerToProperty(r.Context(), owner_id, property_id); err != nil {
		response.End(w, err.Code, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
