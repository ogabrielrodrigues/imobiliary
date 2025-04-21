package owner

import (
	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/address"
	types "github.com/ogabrielrodrigues/imobiliary/internal/types/marital_status"
)

type DTO struct {
	ID            uuid.UUID           `json:"id"`
	Fullname      string              `json:"fullname"`
	CPF           string              `json:"cpf"`
	RG            string              `json:"rg"`
	Email         string              `json:"email"`
	Cellphone     string              `json:"cellphone"`
	Occupation    string              `json:"occupation"`
	MaritalStatus types.MaritalStatus `json:"marital_status"`
	ManagerID     uuid.UUID           `json:"manager_id"`
	Address       address.DTO         `json:"address"`
}

func (dto *DTO) ToOwner() *Owner {
	return &Owner{
		ID:            dto.ID,
		Fullname:      dto.Fullname,
		CPF:           dto.CPF,
		RG:            dto.RG,
		Email:         dto.Email,
		Cellphone:     dto.Cellphone,
		Occupation:    dto.Occupation,
		MaritalStatus: dto.MaritalStatus,
		Address:       address.New(dto.Address.Street, dto.Address.Number, dto.Address.Complement, dto.Address.Neighborhood, dto.Address.City, dto.Address.State, dto.Address.ZipCode),
	}
}

type CreateDTO struct {
	Fullname      string              `json:"fullname"`
	CPF           string              `json:"cpf"`
	RG            string              `json:"rg"`
	Email         string              `json:"email"`
	Cellphone     string              `json:"cellphone"`
	Occupation    string              `json:"occupation"`
	MaritalStatus types.MaritalStatus `json:"marital_status"`
	Address       address.DTO         `json:"address"`
}
