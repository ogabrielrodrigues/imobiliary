package tenant

import (
	"imobiliary/internal/domain/types"

	"github.com/google/uuid"
)

type Tenant struct {
	ID            uuid.UUID
	ManagerID     uuid.UUID
	Fullname      string
	CPF           types.CPF
	RG            types.RG
	Phone         types.Phone
	Occupation    string
	MaritalStatus types.MaritalStatus
	Address       types.Address
}

func NewTenant(
	manager_id uuid.UUID,
	fullname string,
	cpf types.CPF,
	rg types.RG,
	phone types.Phone,
	occupation string,
	marital_status types.MaritalStatus,
	address types.Address,
) *Tenant {
	newTenant := &Tenant{
		ID:            uuid.New(),
		ManagerID:     manager_id,
		Fullname:      fullname,
		CPF:           cpf,
		RG:            rg,
		Phone:         phone,
		Occupation:    occupation,
		MaritalStatus: marital_status,
		Address:       address,
	}

	return newTenant
}
