package property

import (
	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/address"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

type Status string

const (
	StatusAvailable   Status = "Disponível"
	StatusOccupied    Status = "Ocupado"
	StatusUnavailable Status = "Indisponível"
)

type Property struct {
	ID       uuid.UUID
	Address  *address.Address
	Status   Status
	WaterID  string
	EnergyID string
	UserID   uuid.UUID
}

func New(address *address.Address, status Status, water_id, energy_id string, user_id uuid.UUID) (*Property, *response.Err) {
	p := &Property{
		ID:       uuid.New(),
		Address:  address,
		Status:   status,
		WaterID:  water_id,
		EnergyID: energy_id,
		UserID:   user_id,
	}

	err := p.validate()
	if err != nil {
		return nil, response.NewErr(err.Code, err.Message)
	}

	return p, nil
}

func (p *Property) ToDTO() *DTO {
	return &DTO{
		ID:       p.ID.String(),
		Address:  *p.Address.ToDTO(),
		Status:   p.Status,
		WaterID:  p.WaterID,
		EnergyID: p.EnergyID,
		UserID:   p.UserID.String(),
	}
}
