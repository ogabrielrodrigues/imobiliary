package usecase

import (
	"context"
	"imobiliary/internal/api/middleware"
	"imobiliary/internal/application/dto/request"
	"imobiliary/internal/application/httperr"
	"imobiliary/internal/domain/manager"
	"imobiliary/internal/domain/types"
	"imobiliary/internal/mail"

	"github.com/google/uuid"
)

type ResetPasswordManager struct {
	repository manager.Repository
	mail       mail.MailManagerRepository
	jwtSecret  string
}

func NewResetPasswordManager(repository manager.Repository, mail mail.MailManagerRepository, jwtSecret string) *ResetPasswordManager {
	return &ResetPasswordManager{repository, mail, jwtSecret}
}

func (rpm *ResetPasswordManager) Execute(ctx context.Context, dto request.ResetPasswordDTO) *httperr.HttpError {
	managerID, ok := ctx.Value(middleware.ManagerIDKey).(string)
	if !ok {
		return httperr.NewUnauthorizedError(ctx, "not authorized")
	}

	email, eerr := types.NewEmail(dto.Email)
	if eerr != nil {
		return httperr.NewUnprocessableEntityError(ctx, eerr.Error())
	}

	found, err := rpm.repository.FindByID(ctx, uuid.MustParse(managerID))
	if err != nil {
		return err
	}

	if found.Email.Value() != email.Value() {
		return httperr.NewForbiddenError(ctx, "this email doesn't belong to this manager.")
	}

	return rpm.mail.ResetPasswordMessage(found.Email)
}
