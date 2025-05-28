package user

import (
	"context"
	"mime/multipart"

	"imobiliary/internal/response"

	"github.com/google/uuid"
)

type IRepository interface {
	ListAll(ctx context.Context) ([]DTO, *response.Err)
	FindByID(ctx context.Context, id uuid.UUID) (*User, *response.Err)
	Create(ctx context.Context, user *User) (uuid.UUID, *response.Err)
	Authenticate(ctx context.Context, dto *AuthDTO) (uuid.UUID, *response.Err)
	ChangeAvatar(ctx context.Context, avatar_url string) *response.Err
}

type IAvatarStorageRepository interface {
	ChangeAvatar(ctx context.Context, avatar multipart.File, mime string) (string, *response.Err)
}
