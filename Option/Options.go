package Option

import (
	"go-review/Company"
)

func WithCompanyName(name string) Company.Option {
	return func(c *Company.Company) { c.SetName(name) }
}
