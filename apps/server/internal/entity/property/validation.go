package property

import (
	"net/http"

	"imobiliary/internal/response"
)

const (
	ERR_EMPTY_WATER_ID                   = "cód. água do imóvel não pode ser vazio"
	ERR_EMPTY_ENERGY_ID                  = "cód. energia do imóvel não pode ser vazio"
	ERR_EMPTY_ADDRESS                    = "endereço do imóvel não pode ser vazio"
	ERR_PROPERTY_NOT_FOUND_OR_NOT_EXISTS = "imóvel não encontrado ou não existente"
	ERR_PROPERTY_ALREADY_EXISTS          = "imóvel já existe"
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
