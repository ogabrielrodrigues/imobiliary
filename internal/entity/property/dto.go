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

func (d *CreateDTO) ToProperty() *Property {
	return &Property{
		Status:   d.Status,
		Kind:     d.Kind,
		WaterID:  d.WaterID,
		EnergyID: d.EnergyID,
		Address:  d.Address.ToAddress(),
	}
}

type UpdateDTO struct {
	Address  address.DTO `json:"address"`
	Status   Status      `json:"status"`
	WaterID  string      `json:"water_id"`
	EnergyID string      `json:"energy_id"`
}

func (d *UpdateDTO) ToProperty() *Property {
	return &Property{
		Address:  d.Address.ToAddress(),
		Status:   d.Status,
		WaterID:  d.WaterID,
		EnergyID: d.EnergyID,
	}
}
