package manager

import (
	"context"
	"imobiliary/internal/domain/types"
	"imobiliary/internal/response"

	"github.com/google/uuid"
)

type Repository interface {
	FindByID(ctx context.Context, managerID uuid.UUID) (*Manager, *response.Err)
	Create(ctx context.Context, manager *Manager) *response.Err
	Authenticate(ctx context.Context, email *types.Email, password string) (uuid.UUID, *response.Err)
}
