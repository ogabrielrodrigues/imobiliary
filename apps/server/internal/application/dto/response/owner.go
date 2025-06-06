package response

import (
	"imobiliary/internal/domain/types"
)

type OwnerDTO struct {
	ID            string              `json:"id"`
	ManagerID     string              `json:"manager_id"`
	Fullname      string              `json:"fullname"`
	CPF           string              `json:"cpf"`
	RG            string              `json:"rg"`
	Email         string              `json:"email"`
	Phone         string              `json:"phone"`
	Occupation    string              `json:"occupation"`
	MaritalStatus types.MaritalStatus `json:"marital_status"`
	Address       AddressDTO          `json:"address"`
}
