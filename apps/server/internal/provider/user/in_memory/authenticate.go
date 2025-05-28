package provider

import (
	"context"
	"net/http"

	"imobiliary/internal/entity/user"
	"imobiliary/internal/response"

	"github.com/google/uuid"
)

func (r *InMemoryUserRepository) Authenticate(ctx context.Context, dto *user.AuthDTO) (uuid.UUID, *response.Err) {
	for _, usr := range r.users {
		if usr.Email == dto.Email {
			if usr.ComparePwd(dto.Password) {
				return usr.ID, nil
			}

			return uuid.Nil, response.NewErr(http.StatusUnauthorized, user.ERR_PASSWORD_INVALID)
		}
	}

	return uuid.Nil, response.NewErr(http.StatusNotFound, user.ERR_USER_NOT_FOUND_OR_NOT_EXISTS)
}
