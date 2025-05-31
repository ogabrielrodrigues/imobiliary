package usecase

import (
	"context"
	"errors"
	appcontext "imobiliary/internal/application/context"
	"imobiliary/internal/application/dto/request"
	"imobiliary/internal/application/httperr"
	"imobiliary/internal/domain/owner"

	"github.com/google/uuid"
)

type CreateOwner struct {
	repository owner.Repository
	validator  *owner.Validator
}

func NewCreateOwner(repository owner.Repository, validator *owner.Validator) *CreateOwner {
	return &CreateOwner{repository, validator}
}

func (co *CreateOwner) Execute(ctx context.Context, dto request.CreateOwnerDTO, managerID uuid.UUID) *httperr.HttpError {
	currentTime, err := appcontext.ExtractCurrentTimeFromContext(ctx)
	if err != nil {
		return httperr.NewInvalidContextError(ctx, appcontext.CurrentTimeContextKey, err.Error())
	}

	newOwner, err := co.validator.Validate(dto, managerID, currentTime)
	if err != nil {
		var validationErrs *httperr.ValidationErrors
		if errors.As(err, &validationErrs) {
			return httperr.NewValidationError(ctx, validationErrs)
		}
		return err.(*httperr.HttpError)
	}

	if err := co.repository.Create(ctx, newOwner, managerID); err != nil {
		return err
	}

	return nil
}
