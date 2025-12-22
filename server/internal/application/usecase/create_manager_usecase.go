package usecase

import (
	"context"
	"errors"
	appcontext "imobiliary/internal/application/context"
	"imobiliary/internal/application/dto/request"
	"imobiliary/internal/application/httperr"
	"imobiliary/internal/domain/manager"
)

type CreateManager struct {
	repository manager.Repository
	validator  *manager.Validator
}

func NewCreateManager(repository manager.Repository, validator *manager.Validator) *CreateManager {
	return &CreateManager{repository, validator}
}

func (cm *CreateManager) Execute(ctx context.Context, dto request.CreateManagerDTO) *httperr.HttpError {
	currentTime, err := appcontext.ExtractCurrentTimeFromContext(ctx)
	if err != nil {
		return httperr.NewInvalidContextError(ctx, appcontext.CurrentTimeContextKey, err.Error())
	}

	newManager, err := cm.validator.Validate(dto, currentTime)
	if err != nil {
		var validationErrs *httperr.ValidationErrors
		if errors.As(err, &validationErrs) {
			return httperr.NewValidationError(ctx, validationErrs)
		}
		return err.(*httperr.HttpError)
	}

	if err := cm.repository.Create(ctx, newManager); err != nil {
		return err
	}

	return nil
}
