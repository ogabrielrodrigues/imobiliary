package provider

import (
	"context"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/owner"
	"github.com/ogabrielrodrigues/imobiliary/internal/response"
)

func (r *InMemoryOwnerRepository) Create(ctx context.Context, owner owner.Owner) (uuid.UUID, *response.Err) {
	r.owners = append(r.owners, owner)

	return owner.ID, nil
}
