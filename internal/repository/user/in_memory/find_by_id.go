package repository

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/user"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

func (r *InMemoryUserRepository) FindByID(ctx context.Context, id uuid.UUID) (*user.User, *response.Err) {
	for _, user := range r.users {
		if user.ID == id {
			return user, nil
		}
	}

	return nil, response.NewErr(http.StatusNotFound, user.ERR_USER_NOT_FOUND_OR_NOT_EXISTS)
}
