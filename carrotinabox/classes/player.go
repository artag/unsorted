package classes

import "fmt"

type Player struct {
	name       string
	score      int
	isComputer bool
}

func NewPlayer(name string, isComputer bool) *Player {
	return &Player{
		isComputer: isComputer,
		name:       name,
		score:      0,
	}
}

func (p *Player) Name() string {
	return p.name
}

func (p *Player) NameWithScore() string {
	return fmt.Sprintf("%s: %d", p.name, p.score)
}

func (p *Player) AddScore() {
	p.score++
}

func (p *Player) IsComputer() bool {
	return p.isComputer
}

func (p *Player) IsHuman() bool {
	return !p.IsComputer()
}
