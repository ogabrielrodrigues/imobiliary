package maritalstatus

type MaritalStatus string

const (
	MaritalStatusSolteiro     MaritalStatus = "Solteiro(a)"
	MaritalStatusCasado       MaritalStatus = "Casado(a)"
	MaritalStatusAmasiado     MaritalStatus = "Amasiado(a)"
	MaritalStatusDivorciado   MaritalStatus = "Divorciado(a)"
	MaritalStatusUniaoEstavel MaritalStatus = "União Estável"
	MaritalStatusViuvo        MaritalStatus = "Viúvo(a)"
)
