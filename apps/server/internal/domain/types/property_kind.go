package types

import "errors"

type PropertyKind string

const (
	PropertyKindResidential PropertyKind = "Residencial"
	PropertyKindCommercial  PropertyKind = "Comercial"
	PropertyKindIndustrial  PropertyKind = "Industrial"
	PropertyKindTerrain     PropertyKind = "Terreno"
	PropertyKindRural       PropertyKind = "Rural"
)

func NewPropertyKind(kind PropertyKind) error {
	switch kind {
	case PropertyKindResidential:
		return nil
	case PropertyKindCommercial:
		return nil
	case PropertyKindIndustrial:
		return nil
	case PropertyKindTerrain:
		return nil
	case PropertyKindRural:
		return nil
	default:
		return errors.New("invalid property status")
	}
}
