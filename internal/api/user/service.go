package user

import (
	"context"

	"github.com/google/uuid"
)

type Service struct {
	repo IRepository
}

type IService interface {
	FindByID(ctx context.Context, id uuid.UUID) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	Create(ctx context.Context, dto *CreateDTO) (uuid.UUID, error)
	Update(ctx context.Context, dto *UpdateDTO) error
	Delete(ctx context.Context, id uuid.UUID) error
}

func NewService(repo IRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) FindByID(ctx context.Context, id uuid.UUID) (*User, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *Service) FindByEmail(ctx context.Context, email string) (*User, error) {
	return s.repo.FindByEmail(ctx, email)
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
