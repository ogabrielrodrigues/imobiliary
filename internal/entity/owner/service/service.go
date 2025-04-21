package owner

import (
	"context"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/owner"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

type Service struct {
	repo owner.IRepository
}

type IService interface {
	Create(ctx context.Context, dto owner.CreateDTO) (uuid.UUID, *response.Err)
	FindByID(ctx context.Context, owner_id uuid.UUID) (*owner.DTO, *response.Err)
	FindAllByManagerID(ctx context.Context) ([]owner.DTO, *response.Err)
	AssignOwnerToProperty(ctx context.Context, owner_id uuid.UUID, property_id uuid.UUID) *response.Err
}

func NewService(repo owner.IRepository) *Service {
	return &Service{repo}
}
