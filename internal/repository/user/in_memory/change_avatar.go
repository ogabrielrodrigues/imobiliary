package repository

import (
	"context"

	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

func (r *InMemoryUserRepository) ChangeAvatar(ctx context.Context, avatar_url string) *response.Err {
	return nil
}
