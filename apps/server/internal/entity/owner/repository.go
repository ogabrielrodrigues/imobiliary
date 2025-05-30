package owner

import (
	"context"

	"imobiliary/internal/response"

	"github.com/google/uuid"
)

type IRepository interface {
	Create(ctx context.Context, owner Owner) (uuid.UUID, *response.Err)
	FindAllByManagerID(ctx context.Context) ([]DTO, *response.Err)
	FindByID(ctx context.Context, owner_id uuid.UUID) (*DTO, *response.Err)
	AssignOwnerToProperty(ctx context.Context, owner_id uuid.UUID, property_id uuid.UUID) *response.Err
}
