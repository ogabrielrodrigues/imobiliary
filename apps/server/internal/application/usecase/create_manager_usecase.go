package usecase

import (
	"context"
	"imobiliary/internal/application/dto/request"
	"imobiliary/internal/domain/manager"
)

type CreateManager struct {
	repository manager.Repository
}

func NewCreateManager(repository manager.Repository) *CreateManager {
	return &CreateManager{repository}
}

func (cm *CreateManager) Execute(ctx context.Context, dto request.CreateManagerDTO) error { // TODO: place error type
	// TODO: validate fields to create manager
	return nil
}
