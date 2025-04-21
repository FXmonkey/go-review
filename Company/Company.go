package Company

type Company struct {
	name          string
	local         []string
	team          []string
	establishDate string
}

type Option func(*Company)

func NewIdolCompany(opts ...Option) *Company {
	c := &Company{
		name:          "default",
		local:         make([]string, 0),
		team:          make([]string, 0),
		establishDate: "",
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (c *Company) SetName(name string) {
	c.name = name
}

func (c *Company) GetName() string {
	return c.name
}

func (c *Company) GetLocal() []string {
	if c.local == nil {
		return []string{}
	}
	localCopy := make([]string, len(c.local))
	copy(localCopy, c.local)
	return localCopy
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
