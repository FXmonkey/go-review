package Company

type CompanyInterface interface {
	GetName() string
	GetEstablishDate() string
	GetLocal() []string
	GetTeam() []string
	UpdateEstablishDate(date string)
	AddLocal(local string)
	AddTeam(team string)
}
