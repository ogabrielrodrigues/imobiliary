package provider

import (
	"context"

	"imobiliary/internal/entity/owner"
	"imobiliary/internal/response"

	"github.com/google/uuid"
)

func (r *InMemoryOwnerRepository) Create(ctx context.Context, owner owner.Owner) (uuid.UUID, *response.Err) {
	r.owners = append(r.owners, owner)

	return owner.ID, nil
}
