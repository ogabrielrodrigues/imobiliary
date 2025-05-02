package user

import (
	"net/http"

	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

func (h *Handler) GetUserPlan(w http.ResponseWriter, r *http.Request) *response.Err {
	plan, r_err := h.plan_service.GetUserPlan(r.Context())
	if r_err != nil {
		return r_err
	}

	return response.End(w, http.StatusOK, plan)
}
