package provider

import (
	"context"

	"imobiliary/internal/entity/user"
	"imobiliary/internal/response"

	"github.com/google/uuid"
)

func (r *InMemoryUserRepository) Create(ctx context.Context, user *user.User) (uuid.UUID, *response.Err) {
	r.users = append(r.users, user)
	return user.ID, nil
}
