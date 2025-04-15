package address

import "fmt"

type Kind string

const (
	Residential Kind = "residential"
	Comercial   Kind = "comercial"
)

type Address struct {
	FullAddress  string
	Street       string
	Number       string
	Neighborhood string
	Complement   string
	City         string
	State        string
	StateAbbr    string
	ZipCode      string
	Kind         Kind
}

func New(street, number, neighboorhood, complement, city, state, state_abbr, zip_code string, kind Kind) *Address {
	fulladdress := genFullAddress(street, number, neighboorhood, city, state_abbr, zip_code)

	return &Address{
		FullAddress:  fulladdress,
		Street:       street,
		Number:       number,
		Neighborhood: neighboorhood,
		Complement:   complement,
		City:         city,
		State:        state,
		StateAbbr:    state_abbr,
		ZipCode:      zip_code,
		Kind:         kind,
	}
}

func genFullAddress(street, number, neighborhood, city, state_abbr, zip_code string) string {
	return fmt.Sprintf("%s, n° %s, %s, %s/%s, %s", street, number, neighborhood, city, state_abbr, zip_code)
}

type DTO struct {
	FullAddress  string `json:"full_address"`
	Street       string `json:"street"`
	Number       string `json:"number"`
	Neighborhood string `json:"neighborhood"`
	Complement   string `json:"complement"`
	City         string `json:"city"`
	State        string `json:"state"`
	StateAbbr    string `json:"state_abbr"`
	ZipCode      string `json:"zip_code"`
	Kind         Kind   `json:"kind"`
}

func (a *Address) ToDTO() *DTO {
	return &DTO{
		FullAddress:  a.FullAddress,
		Street:       a.Street,
		Number:       a.Number,
		Neighborhood: a.Neighborhood,
		Complement:   a.Complement,
		City:         a.City,
		State:        a.State,
		StateAbbr:    a.StateAbbr,
		ZipCode:      a.ZipCode,
		Kind:         a.Kind,
	}
}

func (a *DTO) ToAddress() *Address {
	return &Address{
		FullAddress:  a.FullAddress,
		Street:       a.Street,
		Number:       a.Number,
		Neighborhood: a.Neighborhood,
		Complement:   a.Complement,
		City:         a.City,
		State:        a.State,
		StateAbbr:    a.StateAbbr,
		ZipCode:      a.ZipCode,
		Kind:         a.Kind,
	}
}
