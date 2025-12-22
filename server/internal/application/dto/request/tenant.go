package request

import "imobiliary/internal/domain/types"

type CreateTenantDTO struct {
	Fullname      string              `json:"fullname"`
	CPF           string              `json:"cpf"`
	RG            string              `json:"rg"`
	Phone         string              `json:"phone"`
	Occupation    string              `json:"occupation"`
	MaritalStatus types.MaritalStatus `json:"marital_status"`
	Address       *CreateAddressDTO   `json:"address"`
}
