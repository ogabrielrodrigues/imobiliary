package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ogabrielrodrigues/imobiliary/internal/entity/user"
	"github.com/ogabrielrodrigues/imobiliary/internal/response"
)

func (h *Handler) Authenticate(w http.ResponseWriter, r *http.Request) *response.Err {
	var dto user.AuthDTO

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		return response.NewErr(http.StatusBadRequest, response.ERR_INVALID_REQUEST_BODY)
	}

	token, err := h.service.Authenticate(r.Context(), &dto)
	if err != nil {
		return err
	}

	w.Header().Add("Authorization", fmt.Sprintf("Bearer %s", token))
	w.WriteHeader(http.StatusOK)
	return nil
}
