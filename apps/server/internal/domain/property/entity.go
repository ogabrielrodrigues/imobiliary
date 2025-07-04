package property

import (
	"imobiliary/internal/domain/types"

	"github.com/google/uuid"
)

type Property struct {
	ID        uuid.UUID
	ManagerID uuid.UUID
	OwnerID   uuid.UUID
	Status    types.PropertyStatus
	Kind      types.PropertyKind
	WaterID   string
	EnergyID  string
	Address   types.Address
}

func NewProperty(
	managerID uuid.UUID,
	ownerID uuid.UUID,
	status types.PropertyStatus,
	kind types.PropertyKind,
	waterID,
	energyID string,
	address types.Address,
) *Property {
	newProperty := &Property{
		ID:        uuid.New(),
		ManagerID: managerID,
		OwnerID:   ownerID,
		Status:    status,
		Kind:      kind,
		WaterID:   waterID,
		EnergyID:  energyID,
		Address:   address,
	}

	return newProperty
}
