package user

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ogabrielrodrigues/imobiliary/internal/entity/plan"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/user"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var dto user.CreateDTO
	ctx := context.Background()

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		err := response.NewErr(http.StatusBadRequest, user.ERR_INVALID_USER_REQUEST_BODY)
		response.End(w, err.Code, err)
		return
	}

	id, err := h.service.Create(ctx, &dto)
	if err != nil {
		response.End(w, err.Code, err)
		return
	}

	plan := plan.New(plan.PlanKindFree, 30, 0, 30)
	err = h.plan_service.AssignPlanToUser(ctx, string(plan.Kind), id, plan)
	if err != nil {
		response.End(w, err.Code, err)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("/users/%s", id))
	w.WriteHeader(http.StatusCreated)
}
