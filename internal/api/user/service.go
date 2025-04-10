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
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, user *User) error
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

func (s *Service) Create(ctx context.Context, user *User) error {
	return s.repo.Create(ctx, user)
}

func (s *Service) Update(ctx context.Context, user *User) error {
	return s.repo.Update(ctx, user)
}

func (s *Service) Delete(ctx context.Context, user *User) error {
	return s.repo.Delete(ctx, user)
}
