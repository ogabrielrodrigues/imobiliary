package usecase

import (
	"context"
	"imobiliary/internal/application/dto/response"
	"imobiliary/internal/application/httperr"
	"imobiliary/internal/domain/property"

	"github.com/google/uuid"
)

type FindAllProperty struct {
	repository property.Repository
}

func NewFindAllProperty(repository property.Repository) *FindAllProperty {
	return &FindAllProperty{repository}
}

func (fa *FindAllProperty) Execute(ctx context.Context, managerID uuid.UUID) ([]response.PropertyDTO, *httperr.HttpError) {
	properties, err := fa.repository.FindAll(ctx, managerID)
	if err != nil {
		return nil, err
	}

	var propertiesDTO []response.PropertyDTO

	for _, property := range properties {
		propertiesDTO = append(propertiesDTO, response.PropertyDTO{
			ID:       property.ID.String(),
			Status:   property.Status,
			Kind:     property.Kind,
			WaterID:  property.WaterID,
			EnergyID: property.EnergyID,
			OwnerID:  property.OwnerID.String(),
			Address: response.AddressDTO{
				FullAddress:  property.Address.FullAddress,
				Street:       property.Address.Street,
				Number:       property.Address.Number,
				Complement:   property.Address.Complement,
				Neighborhood: property.Address.Neighborhood,
				City:         property.Address.City,
				State:        property.Address.State,
				ZipCode:      property.Address.ZipCode,
			},
		})
	}

	return propertiesDTO, nil
}
