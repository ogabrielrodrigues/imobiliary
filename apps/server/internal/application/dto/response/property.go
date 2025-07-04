package response

import "imobiliary/internal/domain/types"

type PropertyDTO struct {
	ID       string               `json:"id"`
	Status   types.PropertyStatus `json:"status"`
	Kind     types.PropertyKind   `json:"kind"`
	WaterID  string               `json:"water_id"`
	EnergyID string               `json:"energy_id"`
	OwnerID  string               `json:"owner_id"`
	Address  AddressDTO           `json:"address"`
}
