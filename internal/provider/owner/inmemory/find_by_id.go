package provider

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/owner"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

func (r *InMemoryOwnerRepository) FindByID(ctx context.Context, owner_id uuid.UUID) (*owner.DTO, *response.Err) {
	for _, o := range r.owners {
		if o.ID == owner_id {
			return o.ToDTO(), nil
		}
	}

	return nil, response.NewErr(http.StatusNotFound, owner.ERR_OWNER_NOT_FOUND_OR_NOT_EXISTS)
}
