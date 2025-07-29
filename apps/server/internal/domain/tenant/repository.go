package tenant

import (
	"context"
	"imobiliary/internal/application/httperr"

	"github.com/google/uuid"
)

type Repository interface {
	FindByID(ctx context.Context, tenantID, managerID uuid.UUID) (*Tenant, *httperr.HttpError)
	Create(ctx context.Context, tenant *Tenant, managerID uuid.UUID) *httperr.HttpError
	FindAll(ctx context.Context, managerID uuid.UUID) ([]Tenant, *httperr.HttpError)
}
