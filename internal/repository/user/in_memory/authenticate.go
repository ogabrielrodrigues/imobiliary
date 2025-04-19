package repository

import (
	"context"
	"net/http"

	"github.com/ogabrielrodrigues/imobiliary/internal/entity/user"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

func (r *InMemoryUserRepository) Authenticate(ctx context.Context, email, password string) (*user.User, *response.Err) {
	for _, usr := range r.users {
		if usr.Email == email {
			if usr.ComparePwd(password) {
				return usr, nil
			}

			return nil, response.NewErr(http.StatusUnauthorized, user.ERR_PASSWORD_INVALID)
		}
	}

	return nil, response.NewErr(http.StatusNotFound, user.ERR_USER_NOT_FOUND_OR_NOT_EXISTS)
}
