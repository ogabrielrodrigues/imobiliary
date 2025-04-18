package user

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

func (h *Handler) UpdateAvatar(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("avatar")
	if err != nil {
		err := response.NewErr(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		response.End(w, err.Code, err)
		return
	}
	defer file.Close()

	r_err := h.service.SaveAvatar(r.Context(), file)
	if r_err != nil {
		response.End(w, r_err.Code, r_err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetUserPlan(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)

	uid, err := uuid.Parse(user_id)
	if err != nil {
		r_err := response.NewErr(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		response.End(w, r_err.Code, r_err)
		return
	}

	plan, r_err := h.plan_service.GetUserPlan(context.Background(), uid)
	if r_err != nil {
		response.End(w, r_err.Code, r_err)
		return
	}

	response.End(w, http.StatusOK, plan)
}
