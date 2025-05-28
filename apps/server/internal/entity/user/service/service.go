package user

import (
	"context"
	"mime/multipart"

	"imobiliary/internal/entity/user"
	"imobiliary/internal/response"

	"github.com/google/uuid"
)

type Service struct {
	repo    user.IRepository
	storage user.IAvatarStorageRepository
}

type IService interface {
	ListAll(ctx context.Context) ([]user.DTO, *response.Err)
	FindByID(ctx context.Context, id uuid.UUID) (*user.DTO, *response.Err)
	Create(ctx context.Context, dto *user.CreateDTO) (uuid.UUID, *response.Err)
	Authenticate(ctx context.Context, dto *user.AuthDTO) (string, *response.Err)
	ChangeAvatar(ctx context.Context, file multipart.File, metadata *multipart.FileHeader) *response.Err
}

func NewService(repo user.IRepository, storage user.IAvatarStorageRepository) *Service {
	return &Service{
		repo:    repo,
		storage: storage,
	}
}
