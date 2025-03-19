package kind

type RentTmplData struct {
	Date    string `json:"date"`
	Rentals []Rent `json:"rentals"`
}

type Rent struct {
	Reference      Reference  `json:"reference"`
	Expiration     Expiration `json:"expiration"`
	Owner          Entity     `json:"owner"`
	Tenant         Entity     `json:"tenant"`
	TenantOther    *Entity    `json:"tenant_other,omitempty"`
	Guarantor      *Entity    `json:"guarantor,omitempty"`
	GuarantorOther *Entity    `json:"guarantor_other,omitempty"`
	Property       Property   `json:"property"`
}

type Reference struct {
	Month string `json:"month"`
	Year  string `json:"year"`
}

type Expiration struct {
	Month string `json:"month"`
	Year  string `json:"year"`
}

type Entity struct {
	Kind          string `json:"kind,omitempty"`
	Name          string `json:"name"`
	Nacionality   string `json:"nacionality"`
	MaritalStatus string `json:"marital_status"`
	Occupation    string `json:"occupation"`
	RG            string `json:"rg"`
	CPF           string `json:"cpf"`
	Address       string `json:"address"`
}

type Property struct {
	Value             string `json:"value"`
	Extense           string `json:"extense"`
	WithoutCommission string `json:"without_commission,omitempty"`
	Kind              string `json:"kind"`
	Address           string `json:"address"`
	PaymentDay        string `json:"payment_day"`
	Guarantee         string `json:"guarantee"`
	SignedIn          string `json:"signed_in"`
	ExpiresIn         string `json:"expires_in"`
	ExtendedIn        string `json:"extended_in"`
	VoucherObs        string `json:"voucher_obs,omitempty"`
	BordereauObs      string `json:"bordereau_obs,omitempty"`
}
