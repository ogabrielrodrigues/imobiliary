package user

import (
	"context"
	"mime/multipart"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/user"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

type Service struct {
	repo    user.IRepository
	storage user.IAvatarStorageRepository
}

type IService interface {
	FindByID(ctx context.Context, id uuid.UUID) (*user.DTO, *response.Err)
	FindByEmail(ctx context.Context, email string) (*user.DTO, *response.Err)
	Create(ctx context.Context, dto *user.CreateDTO) (uuid.UUID, *response.Err)
	Authenticate(ctx context.Context, email, password string) (string, *response.Err)
	SaveAvatar(ctx context.Context, avatarFile multipart.File) *response.Err
}

func NewService(repo user.IRepository, storage user.IAvatarStorageRepository) *Service {
	return &Service{
		repo:    repo,
		storage: storage,
	}
}
