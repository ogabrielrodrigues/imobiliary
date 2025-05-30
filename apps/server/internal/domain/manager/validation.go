package manager

import (
	"context"
	"imobiliary/internal/application/dto/request"
	"imobiliary/internal/application/httperr"
	"imobiliary/internal/domain/types"
	"time"
)

type Validator struct{}

func NewManagerValidator() *Validator {
	return &Validator{}
}

func (v *Validator) Validate(dto request.CreateManagerDTO, currentTime time.Time) (*Manager, error) {
	validationErrs := &httperr.ValidationErrors{}

	if dto.Fullname == "" {
		validationErrs.Add("fullname", dto.Fullname, httperr.Required, "fullname must not be empty")
	}

	if len(dto.Fullname) < 10 {
		validationErrs.Add("fullname", dto.Fullname, httperr.MinLength, "fullname must be at least 10 characters")
	}

	if dto.Phone == "" {
		validationErrs.Add("phone", dto.Phone, httperr.Required, "phone must not be empty")
	}

	if dto.Email == "" {
		validationErrs.Add("email", dto.Email, httperr.Required, "email must not be empty")
	}

	if dto.Password == "" {
		validationErrs.Add("password", dto.Password, httperr.Required, "password must not be empty")
	}

	if len(dto.Password) < 8 {
		validationErrs.Add("password", dto.Password, httperr.MinLength, "password must be at least 8 characters")
	}

	phone, err := types.NewPhone(dto.Phone)
	if err != nil {
		validationErrs.Add("phone", dto.Phone, httperr.InvalidFormat, err.Error())
	}

	email, err := types.NewEmail(dto.Email)
	if err != nil {
		validationErrs.Add("email", dto.Email, httperr.InvalidFormat, err.Error())
	}

	if validationErrs.HasErrors() {
		return nil, validationErrs
	}

	newManager, err := NewManager(
		dto.Fullname,
		*phone,
		*email,
		dto.Password,
	)
	if err.(*httperr.HttpError) != nil {
		return nil, httperr.NewUnprocessableEntityError(context.Background(), err.Error())
	}

	return newManager, nil
}
