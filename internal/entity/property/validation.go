package property

import (
	"net/http"

	"github.com/ogabrielrodrigues/imobiliary/internal/response"
)

const (
	ERR_EMPTY_WATER_ID                   = "water_id is empty"
	ERR_EMPTY_ENERGY_ID                  = "energy_id is empty"
	ERR_EMPTY_ADDRESS                    = "address is empty"
	ERR_PROPERTY_NOT_FOUND_OR_NOT_EXISTS = "property not found or not exists"
)

func (p *Property) validate() *response.Err {
	if p.WaterID == "" {
		return response.NewErr(http.StatusBadRequest, ERR_EMPTY_WATER_ID)
	}

	if p.EnergyID == "" {
		return response.NewErr(http.StatusBadRequest, ERR_EMPTY_ENERGY_ID)
	}

	if p.Address == nil {
		return response.NewErr(http.StatusBadRequest, ERR_EMPTY_ADDRESS)
	}

	return nil
}
