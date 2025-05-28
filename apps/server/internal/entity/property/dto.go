package property

import (
	types "imobiliary/internal/types"

	"github.com/google/uuid"
)

type DTO struct {
	ID       string           `json:"id"`
	Status   Status           `json:"status"`
	Kind     Kind             `json:"kind"`
	WaterID  string           `json:"water_id"`
	EnergyID string           `json:"energy_id"`
	OwnerID  string           `json:"owner_id"`
	Address  types.AddressDTO `json:"address"`
}

type CreateDTO struct {
	Status   Status                  `json:"status"`
	Kind     Kind                    `json:"kind"`
	WaterID  string                  `json:"water_id"`
	EnergyID string                  `json:"energy_id"`
	Address  *types.AddressCreateDTO `json:"address"`
}

func (d *DTO) ToProperty() *Property {
	return &Property{
		ID:       uuid.MustParse(d.ID),
		Status:   d.Status,
		Kind:     d.Kind,
		WaterID:  d.WaterID,
		EnergyID: d.EnergyID,
		OwnerID:  uuid.MustParse(d.OwnerID),
		Address:  d.Address.ToAddress(),
	}
}
