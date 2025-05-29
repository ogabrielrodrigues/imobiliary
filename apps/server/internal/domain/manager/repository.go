package manager

import (
	"context"
	"imobiliary/internal/application/httperr"
	"imobiliary/internal/domain/types"

	"github.com/google/uuid"
)

type Repository interface {
	FindByID(ctx context.Context, managerID uuid.UUID) (*Manager, *httperr.HttpError)
	Create(ctx context.Context, manager *Manager) *httperr.HttpError
	Authenticate(ctx context.Context, email *types.Email, password string) (uuid.UUID, *httperr.HttpError)
}
