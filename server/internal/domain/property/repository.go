package property

import (
	"context"
	"imobiliary/internal/application/httperr"

	"github.com/google/uuid"
)

type Repository interface {
	FindByID(ctx context.Context, propertyID, managerID uuid.UUID) (*Property, *httperr.HttpError)
	Create(ctx context.Context, property *Property, managerID uuid.UUID) *httperr.HttpError
	FindAll(ctx context.Context, managerID uuid.UUID) ([]Property, *httperr.HttpError)
}
