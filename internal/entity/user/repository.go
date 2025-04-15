package user

import (
	"context"
	"mime/multipart"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

type Repository struct{}

type IRepository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*User, *response.Err)
	FindByEmail(ctx context.Context, email string) (*User, *response.Err)
	Create(ctx context.Context, user *User) (uuid.UUID, *response.Err)
	Update(ctx context.Context, user *User) *response.Err
	Delete(ctx context.Context, id uuid.UUID) *response.Err
	Authenticate(ctx context.Context, email, password string) (*User, *response.Err)
}

type IAvatarStorageRepository interface {
	GetAvatar(ctx context.Context, id string) string
	SaveAvatar(ctx context.Context, id string, avatar multipart.File) *response.Err
}

func NewRepository() *Repository {
	return &Repository{}
}
