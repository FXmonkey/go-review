package company_do

import "go-review/Interface-crane/Crane"

type ConstructionCompany struct {
	Name  string
	Crane Crane.Crane
}

func (c *ConstructionCompany) do() {
	c.Crane.Up()
	c.Crane.Down()
}
