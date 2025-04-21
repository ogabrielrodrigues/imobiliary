package owner

import (
	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/address"
	types "github.com/ogabrielrodrigues/imobiliary/internal/types/marital_status"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
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
	Address       *address.Address
}

func New(fullname, cpf, rg, email, cellphone, occupation string, marital_status types.MaritalStatus, address *address.Address) (*Owner, *response.Err) {
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
	}
}
