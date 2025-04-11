package Idol

type Idol struct {
	name    string
	age     int
	team    string
	company string
}

func NewIdol(name string, age int, team string, company string) *Idol {
	return &Idol{
		name:    name,
		age:     age,
		team:    team,
		company: company,
	}
}

func (i *Idol) GetName() string {
	return i.name
}
func (i *Idol) GetAge() int {
	return i.age
}
func (i *Idol) GetTeam() string {
	return i.team
}

func (i *Idol) GetCompany() string {
	return i.company
}
