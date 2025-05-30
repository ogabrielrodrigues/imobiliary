package property

import (
	"context"
	"net/http"

	"imobiliary/internal/entity/property"
	"imobiliary/internal/response"

	"github.com/google/uuid"
)

func (s *Service) Create(ctx context.Context, dto *property.CreateDTO) *response.Err {
	if dto.Address == nil {
		return response.NewErr(http.StatusBadRequest, property.ERR_EMPTY_ADDRESS)
	}

	p, err := property.New(dto.Status, dto.Kind, dto.WaterID, dto.EnergyID, uuid.Nil, dto.Address.ToAddress())
	if err != nil {
		return err
	}

	return s.repo.Create(ctx, p)
}
