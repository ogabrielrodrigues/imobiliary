package user

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ogabrielrodrigues/imobiliary/internal/entity/user"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

func (h *Handler) Authenticate(w http.ResponseWriter, r *http.Request) *response.Err {
	var dto user.AuthDTO
	ctx := context.Background()

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		return response.NewErr(http.StatusBadRequest, user.ERR_INVALID_USER_REQUEST_BODY)
	}

	token, err := h.service.Authenticate(ctx, dto.Email, dto.Password)
	if err != nil {
		return err
	}

	w.Header().Add("Authorization", fmt.Sprintf("Bearer %s", token))
	w.WriteHeader(http.StatusOK)

	return nil
}
