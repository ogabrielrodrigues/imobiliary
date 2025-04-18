package property

import (
	"context"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

type IService interface {
	FindAllByUserID(ctx context.Context, user_id uuid.UUID) ([]DTO, *response.Err)
	FindByID(ctx context.Context, user_id uuid.UUID) (*DTO, *response.Err)
	Create(ctx context.Context, dto *CreateDTO, user_id uuid.UUID) (*Property, *response.Err)
}

type Service struct {
	repo IRepository
}

func NewService(repo IRepository) *Service {
	return &Service{repo}
}

func (s *Service) FindAllByUserID(ctx context.Context, user_id uuid.UUID) ([]DTO, *response.Err) {
	return s.repo.FindAllByUserID(ctx, user_id)
}

func (s *Service) FindByID(ctx context.Context, user_id uuid.UUID) (*DTO, *response.Err) {
	property, err := s.repo.FindByID(ctx, user_id)

	return property.ToDTO(), err
}

func (s *Service) Create(ctx context.Context, dto *CreateDTO, user_id uuid.UUID) (*Property, *response.Err) {
	p, err := New(dto.Address.ToAddress(), dto.Status, dto.WaterID, dto.EnergyID, user_id)
	if err != nil {
		return nil, err
	}

	return s.repo.Create(ctx, p)
}
