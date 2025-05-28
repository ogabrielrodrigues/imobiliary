package types

import (
	"imobiliary/internal/response"
	"net/http"
	"regexp"
)

var (
	phoneRegex = regexp.MustCompile(`^\(\d{2}\)\s\d{4,5}-\d{4}$`)
)

type Phone struct {
	Number string
}

func NewPhone(number string) (*Phone, *response.Err) {
	if !phoneRegex.MatchString(number) {
		return nil, response.NewErr(http.StatusUnprocessableEntity, "invalid phone number")
	}

	return &Phone{Number: number}, nil
}

func (p Phone) Value() string {
	return p.Number
}
