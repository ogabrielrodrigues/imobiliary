package types

import "fmt"

type Address struct {
	FullAddress  string
	MiniAddress  string
	Street       string
	Number       string
	Neighborhood string
	Complement   string
	City         string
	State        string
	ZipCode      string
}

func NewAddress(street, number, neighboorhood, complement, city, state, zip_code string) *Address {
	fulladdress := genFullAddress(street, number, neighboorhood, city, state, zip_code)
	miniaddress := genMiniAddress(street, number, neighboorhood, city, state)

	return &Address{
		FullAddress:  fulladdress,
		MiniAddress:  miniaddress,
		Street:       street,
		Number:       number,
		Neighborhood: neighboorhood,
		Complement:   complement,
		City:         city,
		State:        state,
		ZipCode:      zip_code,
	}
}

func genFullAddress(street, number, neighborhood, city, state, zip_code string) string {
	return fmt.Sprintf("%s, n° %s, %s, %s/%s, %s", street, number, neighborhood, city, state, zip_code)
}

func genMiniAddress(street, number, neighborhood, city, state string) string {
	return fmt.Sprintf("%s, %s, %s, %s/%s", street, number, neighborhood, city, state)
}

type AdressDTO struct {
	FullAddress  string `json:"full_address,omitempty"`
	MiniAddress  string `json:"mini_address"`
	Street       string `json:"street"`
	Number       string `json:"number"`
	Neighborhood string `json:"neighborhood"`
	Complement   string `json:"complement,omitempty"`
	City         string `json:"city"`
	State        string `json:"state"`
	ZipCode      string `json:"zip_code"`
}

type AdressCreateDTO struct {
	Street       string `json:"street"`
	Number       string `json:"number"`
	Neighborhood string `json:"neighborhood"`
	Complement   string `json:"complement"`
	City         string `json:"city"`
	State        string `json:"state"`
	ZipCode      string `json:"zip_code"`
}

func (d *AdressCreateDTO) ToAddress() *Address {
	return NewAddress(d.Street, d.Number, d.Neighborhood, d.Complement, d.City, d.State, d.ZipCode)
}

func (a *Address) ToDTO() *AdressDTO {
	return &AdressDTO{
		FullAddress:  a.FullAddress,
		MiniAddress:  a.MiniAddress,
		Street:       a.Street,
		Number:       a.Number,
		Neighborhood: a.Neighborhood,
		Complement:   a.Complement,
		City:         a.City,
		State:        a.State,
		ZipCode:      a.ZipCode,
	}
}

func (a *AdressDTO) ToAddress() *Address {
	address := NewAddress(
		a.Street,
		a.Number,
		a.Neighborhood,
		a.Complement,
		a.City,
		a.State,
		a.ZipCode,
	)

	return address
}
