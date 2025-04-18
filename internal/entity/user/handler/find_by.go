package user

import (
	"context"
	"net/http"
	"regexp"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/user"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

func (h *Handler) FindBy(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	ctx := context.Background()

	email := params.Get("email")
	id := params.Get("id")

	if email != "" && id != "" {
		err := response.NewErr(http.StatusBadRequest, user.ERR_ONLY_ONE_MUST_PARAMETER_MUST_BE_PROVIDED)
		response.End(w, err.Code, err)
		return
	}

	if id == "" {
		if match, _ := regexp.MatchString(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`, email); !match {
			err := response.NewErr(http.StatusBadRequest, user.ERR_EMAIL_INVALID)
			response.End(w, err.Code, err)
			return
		}

		user, err := h.service.FindByEmail(ctx, email)
		if err != nil {
			response.End(w, err.Code, err)
			return
		}

		response.End(w, http.StatusOK, user)
		return
	}

	if email == "" {
		uid, u_err := uuid.Parse(id)
		if u_err != nil {
			err := response.NewErr(http.StatusBadRequest, user.ERR_UUID_INVALID)
			response.End(w, err.Code, err)
			return
		}

		user, err := h.service.FindByID(ctx, uid)
		if err != nil {
			response.End(w, err.Code, err)
			return
		}

		response.End(w, http.StatusOK, user)
		return
	}
}
