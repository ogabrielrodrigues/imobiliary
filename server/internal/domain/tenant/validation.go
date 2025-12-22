package tenant

import (
	"imobiliary/internal/application/dto/request"
	"imobiliary/internal/application/httperr"
	"imobiliary/internal/domain/types"
	"time"

	"github.com/google/uuid"
)

type Validator struct{}

func NewTenantValidator() *Validator {
	return &Validator{}
}

func (v *Validator) Validate(dto request.CreateTenantDTO, managerID uuid.UUID, currentTime time.Time) (*Tenant, error) {
	validationErrs := &httperr.ValidationErrors{}

	if dto.Fullname == "" {
		validationErrs.Add("fullname", dto.Fullname, httperr.Required, "fullname must not be empty")
	}

	if len(dto.Fullname) < 10 {
		validationErrs.Add("fullname", dto.Fullname, httperr.MinLength, "fullname must be at least 10 characters")
	}

	if dto.CPF == "" {
		validationErrs.Add("cpf", dto.CPF, httperr.Required, "cpf must not be empty")
	}

	if dto.RG == "" {
		validationErrs.Add("rg", dto.RG, httperr.Required, "rg must not be empty")
	}

	if dto.Phone == "" {
		validationErrs.Add("phone", dto.Phone, httperr.Required, "phone must not be empty")
	}

	if dto.Occupation == "" {
		validationErrs.Add("occupation", dto.Occupation, httperr.Required, "occupation must not be empty")
	}

	if dto.MaritalStatus == "" {
		validationErrs.Add("marital_status", dto.MaritalStatus, httperr.Required, "marital_status must not be empty")
	}

	if dto.Address == nil {
		validationErrs.Add("address", dto.Address, httperr.Required, "address must not be empty")
	}

	cpf, err := types.NewCPF(dto.CPF)
	if err != nil {
		validationErrs.Add("cpf", dto.CPF, httperr.InvalidFormat, err.Error())
	}

	rg, err := types.NewRG(dto.RG)
	if err != nil {
		validationErrs.Add("rg", dto.RG, httperr.InvalidFormat, err.Error())
	}

	phone, err := types.NewPhone(dto.Phone)
	if err != nil {
		validationErrs.Add("phone", dto.Phone, httperr.InvalidFormat, err.Error())
	}

	address, err := types.NewAddress(
		dto.Address.Street,
		dto.Address.Number,
		dto.Address.Complement,
		dto.Address.Neighborhood,
		dto.Address.City,
		dto.Address.State,
		dto.Address.ZipCode,
	)
	if err.(*httperr.ValidationErrors) != nil {
		if err.(*httperr.ValidationErrors).HasErrors() {
			validationErrs.Merge(err.(*httperr.ValidationErrors))
		}
	}

	if validationErrs.HasErrors() {
		return nil, validationErrs
	}

	newTenant := NewTenant(
		managerID,
		dto.Fullname,
		*cpf,
		*rg,
		*phone,
		dto.Occupation,
		dto.MaritalStatus,
		*address,
	)

	return newTenant, nil
}
