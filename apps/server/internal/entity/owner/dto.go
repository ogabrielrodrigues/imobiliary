package owner

import (
	"imobiliary/internal/types"

	"github.com/google/uuid"
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
	Address       types.AddressDTO    `json:"address"`
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
		Address:       types.NewAddress(dto.Address.Street, dto.Address.Number, dto.Address.Complement, dto.Address.Neighborhood, dto.Address.City, dto.Address.State, dto.Address.ZipCode),
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
	Address       types.AddressDTO    `json:"address"`
}

type AssignOwnerDTO struct {
	OwnerID    string `json:"owner_id"`
	PropertyID string `json:"property_id"`
}
