package property

import (
	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/address"
)

type DTO struct {
	ID       string      `json:"id"`
	Status   Status      `json:"status"`
	Kind     Kind        `json:"kind"`
	WaterID  string      `json:"water_id"`
	EnergyID string      `json:"energy_id"`
	UserID   string      `json:"user_id"`
	OwnerID  string      `json:"owner_id"`
	Address  address.DTO `json:"address"`
}

func (d *DTO) ToProperty() *Property {
	return &Property{
		ID:       uuid.MustParse(d.ID),
		Status:   d.Status,
		Kind:     d.Kind,
		WaterID:  d.WaterID,
		EnergyID: d.EnergyID,
		UserID:   uuid.MustParse(d.UserID),
		OwnerID:  uuid.MustParse(d.OwnerID),
		Address:  d.Address.ToAddress(),
	}
}

type CreateDTO struct {
	Status   Status            `json:"status"`
	Kind     Kind              `json:"kind"`
	WaterID  string            `json:"water_id"`
	EnergyID string            `json:"energy_id"`
	Address  address.CreateDTO `json:"address"`
}
