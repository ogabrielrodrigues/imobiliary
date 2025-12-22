package types

import (
	"errors"
)

type RG struct {
	Rg string
}

func NewRG(rg string) (*RG, error) {
	if len(rg) < 5 {
		return nil, errors.New("invalid RG")
	}

	return &RG{Rg: rg}, nil
}

func (r RG) Value() string {
	return r.Rg
}
