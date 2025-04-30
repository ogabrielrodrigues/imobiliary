package user

import (
	"context"
	"mime/multipart"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

type IRepository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*User, *response.Err)
	Create(ctx context.Context, user *User) (uuid.UUID, *response.Err)
	Authenticate(ctx context.Context, email, password string) (*User, *response.Err)
	ChangeAvatar(ctx context.Context, avatar_url string) *response.Err
}

type IAvatarStorageRepository interface {
	SaveAvatar(ctx context.Context, avatar multipart.File) (string, *response.Err)
}
