package provider

import (
	"context"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

func (r *InMemoryOwnerRepository) AssignOwnerToProperty(ctx context.Context, owner_id uuid.UUID, property_id uuid.UUID) *response.Err {
	return nil
}
