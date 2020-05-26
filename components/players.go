package components

type Player struct {
	name string
}

func (p *Player) SetName(name string) {
	p.name = name
}

func (p *Player) GetName() string {
	return p.name
}
