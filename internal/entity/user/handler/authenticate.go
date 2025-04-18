package user

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ogabrielrodrigues/imobiliary/internal/entity/user"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

func (h *Handler) Authenticate(w http.ResponseWriter, r *http.Request) {
	var dto user.AuthDTO
	ctx := context.Background()

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		err := response.NewErr(http.StatusBadRequest, user.ERR_INVALID_USER_REQUEST_BODY)
		response.End(w, err.Code, err)
	}

	token, err := h.service.Authenticate(ctx, dto.Email, dto.Password)
	if err != nil {
		if err.Message == user.ERR_USER_NOT_FOUND_OR_NOT_EXISTS {
			response.End(w, err.Code, err)
			return
		}

		response.End(w, err.Code, err)
		return
	}

	w.Header().Add("Authorization", fmt.Sprintf("Bearer %s", token))
	w.WriteHeader(http.StatusOK)
}
