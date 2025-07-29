package usecase

import (
	"context"
	"errors"
	appcontext "imobiliary/internal/application/context"
	"imobiliary/internal/application/dto/request"
	"imobiliary/internal/application/httperr"
	"imobiliary/internal/domain/tenant"

	"github.com/google/uuid"
)

type CreateTenant struct {
	repository tenant.Repository
	validator  *tenant.Validator
}

func NewCreateTenant(repository tenant.Repository, validator *tenant.Validator) *CreateTenant {
	return &CreateTenant{repository, validator}
}

func (ct *CreateTenant) Execute(ctx context.Context, dto request.CreateTenantDTO, managerID uuid.UUID) *httperr.HttpError {
	currentTime, err := appcontext.ExtractCurrentTimeFromContext(ctx)
	if err != nil {
		return httperr.NewInvalidContextError(ctx, appcontext.CurrentTimeContextKey, err.Error())
	}

	newTenant, err := ct.validator.Validate(dto, managerID, currentTime)
	if err != nil {
		var validationErrs *httperr.ValidationErrors
		if errors.As(err, &validationErrs) {
			return httperr.NewValidationError(ctx, validationErrs)
		}
		return err.(*httperr.HttpError)
	}

	if err := ct.repository.Create(ctx, newTenant, managerID); err != nil {
		return err
	}

	return nil
}
