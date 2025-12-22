package usecase

import (
	"context"
	"errors"
	appcontext "imobiliary/internal/application/context"
	"imobiliary/internal/application/dto/request"
	"imobiliary/internal/application/httperr"
	"imobiliary/internal/domain/property"

	"github.com/google/uuid"
)

type CreateProperty struct {
	repository property.Repository
	validator  *property.Validator
}

func NewCreateProperty(repository property.Repository, validator *property.Validator) *CreateProperty {
	return &CreateProperty{repository, validator}
}

func (cp *CreateProperty) Execute(ctx context.Context, dto request.CreatePropertyDTO, managerID uuid.UUID) *httperr.HttpError {
	currentTime, err := appcontext.ExtractCurrentTimeFromContext(ctx)
	if err != nil {
		return httperr.NewInvalidContextError(ctx, appcontext.CurrentTimeContextKey, err.Error())
	}

	newProperty, err := cp.validator.Validate(dto, managerID, currentTime)
	if err != nil {
		var validationErrs *httperr.ValidationErrors
		if errors.As(err, &validationErrs) {
			return httperr.NewValidationError(ctx, validationErrs)
		}
		return err.(*httperr.HttpError)
	}

	if err := cp.repository.Create(ctx, newProperty, managerID); err != nil {
		return err
	}

	return nil
}
