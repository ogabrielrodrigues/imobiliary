package property

import (
	"context"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

type IService interface {
	FindAllByUserID(ctx context.Context, user_id uuid.UUID) ([]Property, *response.Err)
	FindByID(ctx context.Context, user_id uuid.UUID) (*Property, *response.Err)
	Create(ctx context.Context, property *Property) (*Property, *response.Err)
	Update(ctx context.Context, property *Property) (*Property, *response.Err)
	Delete(ctx context.Context, user_id uuid.UUID) *response.Err
}

type Service struct {
	repo IRepository
}

func NewService(repo IRepository) *Service {
	return &Service{repo}
}

func (s *Service) FindAllByUserID(ctx context.Context, user_id uuid.UUID) ([]Property, *response.Err) {
	return s.repo.FindAllByUserID(ctx, user_id)
}

func (s *Service) FindByID(ctx context.Context, user_id uuid.UUID) (*Property, *response.Err) {
	return s.repo.FindByID(ctx, user_id)
}

func (s *Service) Create(ctx context.Context, property *Property) (*Property, *response.Err) {
	return s.repo.Create(ctx, property)
}

func (s *Service) Update(ctx context.Context, property *Property) (*Property, *response.Err) {
	return s.repo.Update(ctx, property)
}

func (s *Service) Delete(ctx context.Context, user_id uuid.UUID) *response.Err {
	return s.repo.Delete(ctx, user_id)
}
