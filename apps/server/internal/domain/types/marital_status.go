package types

type MaritalStatus string

const (
	Single      MaritalStatus = "Solteiro(a)"
	Married     MaritalStatus = "Casado(a)"
	Loved       MaritalStatus = "Amasiado(a)"
	Divorced    MaritalStatus = "Divorciado(a)"
	StableUnion MaritalStatus = "União Estável"
	Widowed     MaritalStatus = "Viúvo(a)"
)
