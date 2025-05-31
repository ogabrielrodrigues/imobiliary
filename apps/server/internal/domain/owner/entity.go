package owner

import (
	"imobiliary/internal/domain/types"

	"github.com/google/uuid"
)

type Owner struct {
	ID            uuid.UUID
	ManagerID     uuid.UUID
	Fullname      string
	CPF           types.CPF
	RG            types.RG
	Phone         types.Phone
	Email         types.Email
	Occupation    string
	MaritalStatus types.MaritalStatus
	Address       types.Address
}

func NewOwner(
	manager_id uuid.UUID,
	fullname string,
	cpf types.CPF,
	rg types.RG,
	phone types.Phone,
	email types.Email,
	occupation string,
	marital_status types.MaritalStatus,
	address types.Address,
) *Owner {
	newOwner := &Owner{
		ID:            uuid.New(),
		ManagerID:     manager_id,
		Fullname:      fullname,
		CPF:           cpf,
		RG:            rg,
		Phone:         phone,
		Email:         email,
		Occupation:    occupation,
		MaritalStatus: marital_status,
		Address:       address,
	}

	return newOwner
}
