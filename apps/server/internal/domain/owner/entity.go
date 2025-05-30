package owner

import (
	"imobiliary/internal/domain/types"

	"github.com/google/uuid"
)

type Owner struct {
	ID            uuid.UUID
	ManagerID     uuid.UUID
	Fullname      string
	CPF           string
	RG            string
	Email         types.Email
	Phone         types.Phone
	Occupation    string
	MaritalStatus types.MaritalStatus
	Address       types.Address
}

func New(
	manager_id uuid.UUID,
	fullname,
	cpf,
	rg string,
	email types.Email,
	phone types.Phone,
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
		Email:         email,
		Phone:         phone,
		Occupation:    occupation,
		MaritalStatus: marital_status,
		Address:       address,
	}

	return newOwner
}
