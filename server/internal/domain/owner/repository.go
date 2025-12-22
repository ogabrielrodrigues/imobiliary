package owner

import (
	"context"
	"imobiliary/internal/application/httperr"

	"github.com/google/uuid"
)

type Repository interface {
	FindByID(ctx context.Context, ownerID, managerID uuid.UUID) (*Owner, *httperr.HttpError)
	Create(ctx context.Context, owner *Owner, managerID uuid.UUID) *httperr.HttpError
	FindAll(ctx context.Context, managerID uuid.UUID) ([]Owner, *httperr.HttpError)
}
