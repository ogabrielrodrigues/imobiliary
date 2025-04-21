package property

import (
	"context"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

type IRepository interface {
	FindAllByUserID(ctx context.Context) ([]DTO, *response.Err)
	FindByID(ctx context.Context, id uuid.UUID) (*DTO, *response.Err)
	Create(ctx context.Context, property *Property) *response.Err
}
