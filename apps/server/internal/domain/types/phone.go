package types

import (
	"errors"
	"regexp"
)

var (
	phoneRegex = regexp.MustCompile(`^\(\d{2}\)\s\d{4,5}-\d{4}$`)
)

type Phone struct {
	Number string
}

func NewPhone(number string) (*Phone, error) {
	if !phoneRegex.MatchString(number) {
		return nil, errors.New("invalid phone number")
	}

	return &Phone{Number: number}, nil
}

func (p Phone) Value() string {
	return p.Number
}
