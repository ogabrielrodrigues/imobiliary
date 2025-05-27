package provider

import (
	"context"

	"github.com/ogabrielrodrigues/imobiliary/internal/response"
)

func (r *InMemoryUserRepository) ChangeAvatar(ctx context.Context, avatar_url string) *response.Err {
	return nil
}
