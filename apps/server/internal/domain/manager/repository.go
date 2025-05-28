package manager

import (
	"context"
	"imobiliary/internal/domain/types"

	"github.com/google/uuid"
)

type Repository interface {
	FindByID(ctx context.Context, managerID uuid.UUID) (*Manager, error)
	Create(ctx context.Context, manager *Manager) error
	Authenticate(ctx context.Context, email *types.Email, password string) (uuid.UUID, error)
}
