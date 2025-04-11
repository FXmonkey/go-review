package Idol

type IdolManager struct {
	idols map[string]*Idol
}

func NewIdolManager() *IdolManager {
	return &IdolManager{
		idols: make(map[string]*Idol),
	}
}

func (im *IdolManager) AddIdol(idol *Idol) {
	im.idols[idol.name] = idol
}

func (im *IdolManager) GetIdol(name string) *Idol {
	idol := im.idols[name]
	return idol
}
