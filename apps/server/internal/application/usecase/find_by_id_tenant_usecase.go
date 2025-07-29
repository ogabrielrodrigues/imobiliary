package usecase

import (
	"context"
	"imobiliary/internal/application/dto/response"
	"imobiliary/internal/application/httperr"
	"imobiliary/internal/domain/tenant"

	"github.com/google/uuid"
)

type FindByIDTenant struct {
	repository tenant.Repository
}

func NewFindByIDTenant(repository tenant.Repository) *FindByIDTenant {
	return &FindByIDTenant{repository}
}

func (ft *FindByIDTenant) Execute(ctx context.Context, tenantID, managerID uuid.UUID) (*response.TenantDTO, *httperr.HttpError) {
	tenant, err := ft.repository.FindByID(ctx, tenantID, managerID)
	if err != nil {
		return nil, err
	}

	tenantDTO := response.TenantDTO{
		ID:            tenant.ID.String(),
		ManagerID:     tenant.ManagerID.String(),
		Fullname:      tenant.Fullname,
		CPF:           tenant.CPF.Value(),
		RG:            tenant.RG.Value(),
		Phone:         tenant.Phone.Value(),
		Occupation:    tenant.Occupation,
		MaritalStatus: tenant.MaritalStatus,
		Address: response.AddressDTO{
			FullAddress:  tenant.Address.FullAddress,
			Street:       tenant.Address.Street,
			Number:       tenant.Address.Number,
			Complement:   tenant.Address.Complement,
			Neighborhood: tenant.Address.Neighborhood,
			City:         tenant.Address.City,
			State:        tenant.Address.State,
			ZipCode:      tenant.Address.ZipCode,
		},
	}

	return &tenantDTO, nil
}
