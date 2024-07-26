package domain

type IDType string

const (
	CEDULA            IDType = "CC"
	TARJETA_IDENTIDAD IDType = "TI"
	PASAPORTE         IDType = "PA"
)

type Person struct {
	Type           IDType
	Identification string
	Name           string
	Age            int8
	Company        Company
	Skills         []Skill
}
