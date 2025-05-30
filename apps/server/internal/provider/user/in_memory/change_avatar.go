package provider

import (
	"context"

	"imobiliary/internal/response"
)

func (r *InMemoryUserRepository) ChangeAvatar(ctx context.Context, avatar_url string) *response.Err {
	return nil
}
