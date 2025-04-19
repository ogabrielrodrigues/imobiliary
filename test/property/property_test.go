package test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/property"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/address"
)

func TestProperty(t *testing.T) {
	t.Run("should be able to create a property", func(t *testing.T) {
		_, err := property.New(
			property.StatusAvailable,
			property.KindResidential,
			"water-123",
			"energy-123",
			uuid.New(),
			address.New("street", "number", "neighborhood", "complement", "city", "state", "zip"),
		)

		if err != nil {
			t.Errorf("expected property to be created, got nil")
		}
	})

	t.Run("should not be able to create a property with empty water_id", func(t *testing.T) {
		_, err := property.New(
			property.StatusAvailable,
			property.KindResidential,
			"",
			"energy-123",
			uuid.New(),
			address.New("street", "number", "neighborhood", "complement", "city", "state", "zip"),
		)

		if err.Message != property.ERR_EMPTY_WATER_ID {
			t.Errorf("expected err: %s, got: %s", property.ERR_EMPTY_WATER_ID, err.Message)
		}
	})

	t.Run("should not be able to create a property with empty energy_id", func(t *testing.T) {
		_, err := property.New(
			property.StatusAvailable,
			property.KindResidential,
			"water-123",
			"",
			uuid.New(),
			address.New("street", "number", "neighborhood", "complement", "city", "state", "zip"),
		)

		if err.Message != property.ERR_EMPTY_ENERGY_ID {
			t.Errorf("expected err: %s, got: %s", property.ERR_EMPTY_ENERGY_ID, err.Message)
		}
	})

	t.Run("should not be able to create a property with empty address", func(t *testing.T) {
		_, err := property.New(
			property.StatusAvailable,
			property.KindResidential,
			"water-123",
			"energy-123",
			uuid.New(),
			nil,
		)

		if err.Message != property.ERR_EMPTY_ADDRESS {
			t.Errorf("expected err: %s, got: %s", property.ERR_EMPTY_ADDRESS, err.Message)
		}
	})
}
