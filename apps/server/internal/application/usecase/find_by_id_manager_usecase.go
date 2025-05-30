package usecase

import (
	"context"
	"imobiliary/internal/application/dto/response"
	"imobiliary/internal/application/httperr"
	"imobiliary/internal/domain/manager"

	"github.com/google/uuid"
)

type FindByIDManager struct {
	repository manager.Repository
}

func NewFindByIDManager(repository manager.Repository) *FindByIDManager {
	return &FindByIDManager{repository}
}

func (cm *FindByIDManager) Execute(ctx context.Context, managerID uuid.UUID) (*response.ManagerDTO, *httperr.HttpError) {
	manager, err := cm.repository.FindByID(ctx, managerID)
	if err != nil {
		return nil, err
	}

	managerDTO := response.ManagerDTO{
		ID:       manager.ID.String(),
		Fullname: manager.Fullname,
		Email:    manager.Email.Value(),
		Phone:    manager.Phone.Value(),
	}

	return &managerDTO, nil
}
