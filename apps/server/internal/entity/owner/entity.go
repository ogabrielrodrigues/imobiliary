package owner

import (
	"imobiliary/internal/response"
	"imobiliary/internal/types"

	"github.com/google/uuid"
)

type Owner struct {
	ID            uuid.UUID
	Fullname      string
	CPF           string
	RG            string
	Email         string
	Cellphone     string
	Occupation    string
	MaritalStatus types.MaritalStatus
	Address       *types.Address
	ManagerID     uuid.UUID
}

func New(fullname, cpf, rg, email, cellphone, occupation string, marital_status types.MaritalStatus, address *types.Address, manager_id uuid.UUID) (*Owner, *response.Err) {
	o := &Owner{
		ID:            uuid.New(),
		Fullname:      fullname,
		CPF:           cpf,
		RG:            rg,
		Email:         email,
		Cellphone:     cellphone,
		Occupation:    occupation,
		MaritalStatus: marital_status,
		Address:       address,
		ManagerID:     manager_id,
	}

	if err := o.validate(); err != nil {
		return nil, err
	}

	return o, nil
}

func (u *Owner) ToDTO() *DTO {
	return &DTO{
		ID:            u.ID,
		Fullname:      u.Fullname,
		CPF:           u.CPF,
		RG:            u.RG,
		Email:         u.Email,
		Cellphone:     u.Cellphone,
		Occupation:    u.Occupation,
		MaritalStatus: u.MaritalStatus,
		Address:       *u.Address.ToDTO(),
		ManagerID:     u.ManagerID,
	}
}
