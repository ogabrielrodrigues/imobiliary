package property

import (
	"context"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/property"
	"github.com/ogabrielrodrigues/imobiliary/internal/response"
)

type IService interface {
	FindAllByUserID(ctx context.Context) ([]property.DTO, *response.Err)
	FindByID(ctx context.Context, id uuid.UUID) (*property.DTO, *response.Err)
	Create(ctx context.Context, dto *property.CreateDTO) *response.Err
}

type Service struct {
	repo property.IRepository
}

func NewService(repo property.IRepository) *Service {
	return &Service{repo}
}
