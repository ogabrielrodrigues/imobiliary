package property

import (
	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/address"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

type Property struct {
	ID       uuid.UUID
	Address  *address.Address
	WaterID  string
	EnergyID string
}

func New(address *address.Address, water_id, energy_id string) (*Property, *response.Err) {
	p := &Property{
		ID:       uuid.New(),
		Address:  address,
		WaterID:  water_id,
		EnergyID: energy_id,
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
		WaterID:  p.WaterID,
		EnergyID: p.EnergyID,
	}
}
