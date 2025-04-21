package provider

import (
	"context"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/user"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

func (r *InMemoryUserRepository) Create(ctx context.Context, user *user.User) (uuid.UUID, *response.Err) {
	r.users = append(r.users, user)
	return user.ID, nil
}
