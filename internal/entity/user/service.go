package user

import (
	"context"
	"mime/multipart"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

type Service struct {
	repo    IRepository
	storage IAvatarStorageRepository
}

type IService interface {
	FindByID(ctx context.Context, id uuid.UUID) (*DTO, *response.Err)
	FindByEmail(ctx context.Context, email string) (*DTO, *response.Err)
	Create(ctx context.Context, dto *CreateDTO) (uuid.UUID, *response.Err)
	Update(ctx context.Context, dto *UpdateDTO) *response.Err
	Delete(ctx context.Context, id uuid.UUID) *response.Err
	Authenticate(ctx context.Context, email, password string) (string, *response.Err)
	SaveAvatar(ctx context.Context, id uuid.UUID, avatarFile multipart.File) *response.Err
}

func NewService(repo IRepository, storage IAvatarStorageRepository) *Service {
	return &Service{
		repo:    repo,
		storage: storage,
	}
}

func (s *Service) FindByID(ctx context.Context, id uuid.UUID) (*DTO, *response.Err) {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user.ToDTO(), nil
}

func (s *Service) FindByEmail(ctx context.Context, email string) (*DTO, *response.Err) {
	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return user.ToDTO(), nil
}

func (s *Service) Create(ctx context.Context, dto *CreateDTO) (uuid.UUID, *response.Err) {
	user, err := New(dto.CreciID, dto.Fullname, dto.Cellphone, dto.Email, dto.Password)
	if err != nil {
		return uuid.Nil, err
	}

	return s.repo.Create(ctx, user)
}

func (s *Service) Update(ctx context.Context, dto *UpdateDTO) *response.Err {
	user := dto.ToUser()

	return s.repo.Update(ctx, user)
}

func (s *Service) Delete(ctx context.Context, id uuid.UUID) *response.Err {
	return s.repo.Delete(ctx, id)
}

func (s *Service) Authenticate(ctx context.Context, email, password string) (string, *response.Err) {
	user, err := s.repo.Authenticate(ctx, email, password)
	if err != nil {
		return "", err
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.ID,
		"sub":  user.ID,
		"user": user.ToDTO(),
		"exp":  time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	token, t_err := claims.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if t_err != nil {
		return "", response.NewErr(http.StatusInternalServerError, ERR_FAILED_GENERATE_TOKEN)
	}

	return token, nil
}

func (s *Service) SaveAvatar(ctx context.Context, id uuid.UUID, avatarFile multipart.File) *response.Err {
	return s.storage.SaveAvatar(ctx, id.String(), avatarFile)
}
