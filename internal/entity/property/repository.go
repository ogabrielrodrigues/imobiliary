package property

import (
	"context"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

type IRepository interface {
	FindAllByUserID(ctx context.Context, user_id uuid.UUID) ([]DTO, *response.Err)
	FindByID(ctx context.Context, id uuid.UUID) (*Property, *response.Err)
	Create(ctx context.Context, property *Property) (*Property, *response.Err)
	Update(ctx context.Context, property *Property) (*Property, *response.Err)
	Delete(ctx context.Context, id uuid.UUID) *response.Err
}

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}
