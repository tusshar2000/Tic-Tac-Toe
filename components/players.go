package components

type Player struct {
	name string
	mark string
}

func (p *Player) SetName(name string) {
	p.name = name
}
func (p *Player) SetMark(mark string) {
	p.mark = mark
}

func (p *Player) GetName() string {
	return p.name
}

func (p *Player) GetMark() string {
	return p.mark
}
