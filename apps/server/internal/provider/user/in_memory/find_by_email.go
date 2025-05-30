package provider

import (
	"context"
	"net/http"

	"imobiliary/internal/entity/user"
	"imobiliary/internal/response"
)

func (r *InMemoryUserRepository) FindByEmail(ctx context.Context, email string) (*user.User, *response.Err) {
	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, response.NewErr(http.StatusNotFound, user.ERR_USER_NOT_FOUND_OR_NOT_EXISTS)
}
