package provider

import (
	"context"
	"net/http"

	"imobiliary/internal/entity/user"
	"imobiliary/internal/response"

	"github.com/google/uuid"
)

func (r *InMemoryUserRepository) FindByID(ctx context.Context, id uuid.UUID) (*user.User, *response.Err) {
	for _, user := range r.users {
		if user.ID == id {
			return user, nil
		}
	}

	return nil, response.NewErr(http.StatusNotFound, user.ERR_USER_NOT_FOUND_OR_NOT_EXISTS)
}
