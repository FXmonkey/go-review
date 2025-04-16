package Company

type CompanyManager struct {
	CompanyList map[string]*Company
}

func NewCompanyManager() *CompanyManager {
	return &CompanyManager{
		CompanyList: make(map[string]*Company),
	}
}

func (cm *CompanyManager) AddCompany(c *Company) {
	cm.CompanyList[c.name] = c
}

func (cm *CompanyManager) GetCompany(name string) (*Company, bool) {
	c, ok := cm.CompanyList[name]
	return c, ok
}

func (cm *CompanyManager) GetAllCompanies() map[string]*Company {
	return cm.CompanyList
}
