package Company

type CompanyManager struct {
	CompanyList map[string]CompanyInterface
}

func NewCompanyManager() *CompanyManager {
	return &CompanyManager{
		CompanyList: make(map[string]CompanyInterface),
	}
}

func (cm *CompanyManager) AddCompany(c CompanyInterface) {
	cm.CompanyList[c.GetName()] = c
}

func (cm *CompanyManager) GetCompany(name string) (CompanyInterface, bool) {
	c, ok := cm.CompanyList[name]
	return c, ok
}

func (cm *CompanyManager) GetAllCompanies() map[string]CompanyInterface {
	return cm.CompanyList
}
