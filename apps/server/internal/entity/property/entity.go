package property

import (
	"imobiliary/internal/response"
	types "imobiliary/internal/types"

	"github.com/google/uuid"
)

type Status string

const (
	StatusAvailable   Status = "Disponível"
	StatusOccupied    Status = "Ocupado"
	StatusUnavailable Status = "Indisponível"
	StatusReserved    Status = "Reservado"
	StatusRenovating  Status = "Reformando"
)

type Kind string

const (
	KindResidential Kind = "Residencial"
	KindCommercial  Kind = "Comercial"
	KindIndustrial  Kind = "Industrial"
	KindTerrain     Kind = "Terreno"
	KindRural       Kind = "Rural"
)

type Property struct {
	ID       uuid.UUID
	Status   Status
	Kind     Kind
	WaterID  string
	EnergyID string
	UserID   uuid.UUID
	OwnerID  uuid.UUID
	Address  *types.Address
}

func New(status Status, kind Kind, water_id, energy_id string, user_id uuid.UUID, address *types.Address) (*Property, *response.Err) {
	p := &Property{
		ID:       uuid.New(),
		Status:   status,
		Kind:     kind,
		WaterID:  water_id,
		EnergyID: energy_id,
		UserID:   user_id,
		Address:  address,
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
		Kind:     p.Kind,
		Status:   p.Status,
		WaterID:  p.WaterID,
		EnergyID: p.EnergyID,
		OwnerID:  p.OwnerID.String(),
		Address:  *p.Address.ToDTO(),
	}
}
