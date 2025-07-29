package usecase

import (
	"context"
	"imobiliary/internal/application/dto/response"
	"imobiliary/internal/application/httperr"
	"imobiliary/internal/domain/tenant"

	"github.com/google/uuid"
)

type FindAllTenant struct {
	repository tenant.Repository
}

func NewFindAllTenant(repository tenant.Repository) *FindAllTenant {
	return &FindAllTenant{repository}
}

func (fa *FindAllTenant) Execute(ctx context.Context, managerID uuid.UUID) ([]response.TenantDTO, *httperr.HttpError) {
	tenants, err := fa.repository.FindAll(ctx, managerID)
	if err != nil {
		return nil, err
	}

	var tenantsDTO []response.TenantDTO

	for _, tenant := range tenants {
		tenantsDTO = append(tenantsDTO, response.TenantDTO{
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
		})
	}

	return tenantsDTO, nil
}
