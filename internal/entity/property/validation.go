package property

import (
	"net/http"

	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
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
