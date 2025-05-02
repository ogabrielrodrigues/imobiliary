package user

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ogabrielrodrigues/imobiliary/internal/entity/user"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) *response.Err {
	var dto user.CreateDTO
	ctx := context.Background()

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		return response.NewErr(http.StatusBadRequest, user.ERR_INVALID_USER_REQUEST_BODY)
	}

	id, err := h.service.Create(ctx, &dto)
	if err != nil {
		return err
	}

	w.Header().Set("Location", fmt.Sprintf("/users/%s", id))
	w.WriteHeader(http.StatusCreated)

	return nil
}
