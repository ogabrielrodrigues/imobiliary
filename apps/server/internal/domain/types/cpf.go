package types

import (
	"errors"
	"regexp"
)

var (
	cpfRegex = regexp.MustCompile(`^\d{3}\.\d{3}\.\d{3}-\d{2}$`)
)

type CPF struct {
	Cpf string
}

func NewCPF(cpf string) (*CPF, error) {
	if !cpfRegex.MatchString(cpf) {
		return nil, errors.New("invalid CPF")
	}

	return &CPF{Cpf: cpf}, nil
}

func (c CPF) Value() string {
	return c.Cpf
}
