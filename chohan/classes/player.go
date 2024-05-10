package classes

type Player struct {
	isComputer bool
	money      int
}

func NewPlayer(money int, isComputer bool) *Player {
	return &Player{isComputer: isComputer, money: money}
}

func (p *Player) AddMoney(money int) {
	p.money += money
}

func (p *Player) RemoveMoney(money int) {
	p.money -= money
}

func (p *Player) GetMoney() int {
	return p.money
}

func (p *Player) CanBet(bet int) bool {
	return p.money-bet >= 0
}
