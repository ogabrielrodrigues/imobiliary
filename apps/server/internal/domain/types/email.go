package types

import (
	"imobiliary/internal/response"
	"net/http"
	"regexp"
	"strings"
)

var (
	emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
)

type Email struct {
	Address string
}

func NewEmail(address string) (*Email, *response.Err) {
	address = strings.TrimSpace(address)

	if !emailRegex.MatchString(address) {
		return nil, response.NewErr(http.StatusUnprocessableEntity, "invalid email address")
	}

	return &Email{Address: address}, nil
}

func (e Email) Value() string {
	return e.Address
}
