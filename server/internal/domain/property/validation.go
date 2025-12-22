package property

import (
	"imobiliary/internal/application/dto/request"
	"imobiliary/internal/application/httperr"
	"imobiliary/internal/domain/types"
	"time"

	"github.com/google/uuid"
)

type Validator struct{}

func NewPropertyValidator() *Validator {
	return &Validator{}
}

func (v *Validator) Validate(dto request.CreatePropertyDTO, managerID uuid.UUID, currentTime time.Time) (*Property, error) {
	validationErrs := &httperr.ValidationErrors{}

	err := types.NewPropertyStatus(dto.Status)
	if err != nil {
		validationErrs.Add("status", dto.Status, httperr.Required, "property status is invalid")
	}

	err = types.NewPropertyKind(dto.Kind)
	if err != nil {
		validationErrs.Add("kind", dto.Kind, httperr.Required, "property kind is invalid")
	}

	if dto.WaterID == "" {
		validationErrs.Add("water_id", dto.WaterID, httperr.Required, "water_id must not be empty")
	}

	if dto.EnergyID == "" {
		validationErrs.Add("energy_id", dto.EnergyID, httperr.Required, "energy_id must not be empty")
	}

	address, err := types.NewAddress(
		dto.Address.Street,
		dto.Address.Number,
		dto.Address.Complement,
		dto.Address.Neighborhood,
		dto.Address.City,
		dto.Address.State,
		dto.Address.ZipCode,
	)
	if err.(*httperr.ValidationErrors) != nil {
		if err.(*httperr.ValidationErrors).HasErrors() {
			validationErrs.Merge(err.(*httperr.ValidationErrors))
		}
	}

	if validationErrs.HasErrors() {
		return nil, validationErrs
	}

	newProperty := NewProperty(
		managerID,
		dto.OwnerID,
		dto.Status,
		dto.Kind,
		dto.WaterID,
		dto.EnergyID,
		*address,
	)

	return newProperty, nil
}
