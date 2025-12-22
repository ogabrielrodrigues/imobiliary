package types

import "errors"

type PropertyStatus string

const (
	PropertyStatusAvailable   PropertyStatus = "Disponível"
	PropertyStatusOccupied    PropertyStatus = "Ocupado"
	PropertyStatusUnavailable PropertyStatus = "Indisponível"
	PropertyStatusReserved    PropertyStatus = "Reservado"
	PropertyStatusRenovating  PropertyStatus = "Reformando"
)

func NewPropertyStatus(status PropertyStatus) error {
	switch status {
	case PropertyStatusAvailable:
		return nil
	case PropertyStatusOccupied:
		return nil
	case PropertyStatusUnavailable:
		return nil
	case PropertyStatusReserved:
		return nil
	case PropertyStatusRenovating:
		return nil
	default:
		return errors.New("invalid property status")
	}
}
