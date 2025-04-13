package user

import (
	"context"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Service struct {
	repo IRepository
}

type IService interface {
	FindByID(ctx context.Context, id uuid.UUID) (*DTO, error)
	FindByEmail(ctx context.Context, email string) (*DTO, error)
	Create(ctx context.Context, dto *CreateDTO) (uuid.UUID, error)
	Update(ctx context.Context, dto *UpdateDTO) error
	Delete(ctx context.Context, id uuid.UUID) error
	Authenticate(ctx context.Context, email, password string) (string, error)
}

func NewService(repo IRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) FindByID(ctx context.Context, id uuid.UUID) (*DTO, error) {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user.ToDTO(), nil
}

func (s *Service) FindByEmail(ctx context.Context, email string) (*DTO, error) {
	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return user.ToDTO(), nil
}

func (s *Service) Create(ctx context.Context, dto *CreateDTO) (uuid.UUID, error) {
	user, err := New(dto.CreciID, dto.Fullname, dto.Cellphone, dto.Email, dto.Password, "")
	if err != nil {
		return uuid.Nil, err
	}

	return s.repo.Create(ctx, user)
}

func (s *Service) Update(ctx context.Context, dto *UpdateDTO) error {
	user := dto.ToUser()

	return s.repo.Update(ctx, user)
}

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

func (s *Service) Authenticate(ctx context.Context, email, password string) (string, error) {
	user, err := s.repo.Authenticate(ctx, email, password)
	if err != nil {
		return "", err
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.ID,
		"user": user.ToDTO(),
		"exp":  time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	token, err := claims.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return token, nil
}
