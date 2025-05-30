package property

import (
	"context"

	"imobiliary/internal/response"

	"github.com/google/uuid"
)

type IRepository interface {
	FindAllByUserID(ctx context.Context) ([]DTO, *response.Err)
	FindByID(ctx context.Context, id uuid.UUID) (*DTO, *response.Err)
	Create(ctx context.Context, property *Property) *response.Err
}
