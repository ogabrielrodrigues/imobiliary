package property

import (
	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/address"
)

type DTO struct {
	ID       string      `json:"id"`
	Address  address.DTO `json:"address"`
	WaterID  string      `json:"water_id"`
	EnergyID string      `json:"energy_id"`
}

func (d *DTO) ToProperty() *Property {
	return &Property{
		ID:       uuid.MustParse(d.ID),
		Address:  d.Address.ToAddress(),
		WaterID:  d.WaterID,
		EnergyID: d.EnergyID,
	}
}
