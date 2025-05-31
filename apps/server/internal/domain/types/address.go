package types

import (
	"fmt"
	"imobiliary/internal/application/httperr"
)

type Address struct {
	FullAddress  string
	Street       string
	Number       string
	Complement   string
	Neighborhood string
	City         string
	State        string
	ZipCode      string
}

func NewAddress(
	street,
	number,
	complement,
	neighborhood,
	city,
	state,
	zip_code string,
) (*Address, *httperr.ValidationErrors) {
	newAddress := &Address{
		Street:       street,
		Number:       number,
		Complement:   complement,
		Neighborhood: neighborhood,
		City:         city,
		State:        state,
		ZipCode:      zip_code,
	}

	newAddress.generateFullAddress()

	if validationErrs := newAddress.validation(); validationErrs.HasErrors() {
		return nil, validationErrs
	}

	return newAddress, nil
}

func (a *Address) generateFullAddress() {
	if a.Complement != "" {
		a.FullAddress = fmt.Sprintf("%s, %s, %s, %s, %s/%s, %s",
			a.Street,
			a.Number,
			a.Complement,
			a.Neighborhood,
			a.City,
			a.State,
			a.ZipCode,
		)
	}

	a.FullAddress = fmt.Sprintf("%s, %s, %s, %s/%s, %s",
		a.Street,
		a.Number,
		a.Neighborhood,
		a.City,
		a.State,
		a.ZipCode,
	)
}

func (a *Address) validation() *httperr.ValidationErrors {
	validationErrs := &httperr.ValidationErrors{}

	if a.Street == "" {
		validationErrs.Add("street", a.Street, httperr.Required, "street must not be empty")
	}

	if a.Number == "" {
		validationErrs.Add("number", a.Number, httperr.Required, "number must not be empty")
	}

	if a.Neighborhood == "" {
		validationErrs.Add("neighborhood", a.Neighborhood, httperr.Required, "neighborhood must not be empty")
	}

	if a.City == "" {
		validationErrs.Add("city", a.City, httperr.Required, "city must not be empty")
	}

	if a.State == "" {
		validationErrs.Add("state", a.State, httperr.Required, "state must not be empty")
	}

	if a.ZipCode == "" {
		validationErrs.Add("zip_code", a.ZipCode, httperr.Required, "zip code must not be empty")
	}

	return validationErrs
}
