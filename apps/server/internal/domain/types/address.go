package types

import "fmt"

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
	neighboorhood,
	city,
	state,
	zip_code string,
) *Address {
	newAddress := &Address{
		Street:       street,
		Number:       number,
		Complement:   complement,
		Neighborhood: neighboorhood,
		City:         city,
		State:        state,
		ZipCode:      zip_code,
	}

	newAddress.generateFullAddress()

	return newAddress
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
