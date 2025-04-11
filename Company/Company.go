package Company

type Company struct {
	name          string
	local         []string
	team          []string
	establishDate string
}

func NewIdolCompany(name string) *Company {
	return &Company{
		name:          name,
		local:         make([]string, 0),
		team:          make([]string, 0),
		establishDate: "",
	}
}

func (c *Company) GetName() string {
	return c.name
}

func (c *Company) GetLocal() []string {
	return c.local
}

func (c *Company) GetTeam() []string {
	return c.team
}

func (c *Company) UpdateName(name string) {
	c.name = name
}

func (c *Company) AddLocal(local string) {
	c.local = append(c.local, local)
}

func (c *Company) AddTeam(team string) {
	c.team = append(c.team, team)
}

func (c *Company) UpdateEstablishDate(date string) {
	c.establishDate = date
}

func (c *Company) GetEstablishDate() string {
	return c.establishDate
}
