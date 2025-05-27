package types

type MaritalStatus string

const (
	MaritalStatusSingle      MaritalStatus = "Solteiro(a)"
	MaritalStatusMarried     MaritalStatus = "Casado(a)"
	MaritalStatusLoved       MaritalStatus = "Amasiado(a)"
	MaritalStatusDivorced    MaritalStatus = "Divorciado(a)"
	MaritalStatusStableUnion MaritalStatus = "União Estável"
	MaritalStatusWidowed     MaritalStatus = "Viúvo(a)"
)
