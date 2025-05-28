package provider

import (
	"context"

	"imobiliary/internal/entity/user"
	"imobiliary/internal/response"
)

func (r *InMemoryUserRepository) ListAll(ctx context.Context) ([]user.DTO, *response.Err) {
	users := make([]user.DTO, 0, len(r.users))

	for _, u := range r.users {
		users = append(users, *u.ToDTO())
	}

	return users, nil
}
