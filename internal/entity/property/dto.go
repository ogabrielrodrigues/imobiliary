package property

import (
	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/address"
)

type DTO struct {
	ID       string      `json:"id"`
	Address  address.DTO `json:"address"`
	Status   Status      `json:"status"`
	WaterID  string      `json:"water_id"`
	EnergyID string      `json:"energy_id"`
	UserID   string      `json:"user_id"`
}

func (d *DTO) ToProperty() *Property {
	return &Property{
		ID:       uuid.MustParse(d.ID),
		Address:  d.Address.ToAddress(),
		Status:   d.Status,
		WaterID:  d.WaterID,
		EnergyID: d.EnergyID,
		UserID:   uuid.MustParse(d.UserID),
	}
}

type CreateDTO struct {
	Address  address.CreateDTO `json:"address"`
	Status   Status            `json:"status"`
	WaterID  string            `json:"water_id"`
	EnergyID string            `json:"energy_id"`
}

func (d *CreateDTO) ToProperty() *Property {
	return &Property{
		Address:  d.Address.ToAddress(),
		Status:   d.Status,
		WaterID:  d.WaterID,
		EnergyID: d.EnergyID,
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
