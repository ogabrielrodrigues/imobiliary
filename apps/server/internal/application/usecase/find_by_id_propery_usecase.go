package usecase

import (
	"context"
	"imobiliary/internal/application/dto/response"
	"imobiliary/internal/application/httperr"
	"imobiliary/internal/domain/property"

	"github.com/google/uuid"
)

type FindByIDProperty struct {
	repository property.Repository
}

func NewFindByIDProperty(repository property.Repository) *FindByIDProperty {
	return &FindByIDProperty{repository}
}

func (fp *FindByIDProperty) Execute(ctx context.Context, propertyID, managerID uuid.UUID) (*response.PropertyDTO, *httperr.HttpError) {
	property, err := fp.repository.FindByID(ctx, propertyID, managerID)
	if err != nil {
		return nil, err
	}

	propertyDTO := response.PropertyDTO{
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
	}

	return &propertyDTO, nil
}
