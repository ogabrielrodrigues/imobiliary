package request

import (
	"imobiliary/internal/domain/types"

	"github.com/google/uuid"
)

type CreatePropertyDTO struct {
	OwnerID  uuid.UUID            `json:"owner_id"`
	Status   types.PropertyStatus `json:"status"`
	Kind     types.PropertyKind   `json:"kind"`
	WaterID  string               `json:"water_id"`
	EnergyID string               `json:"energy_id"`
	Address  *CreateAddressDTO    `json:"address"`
}
