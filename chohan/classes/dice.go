package classes

import (
	"fmt"
	"io"
	"math/rand"
)

type Dice int

type Dices struct {
	dice1 Dice
	dice2 Dice
}

const (
	ICHI Dice = 1
	NI   Dice = 2
	SAN  Dice = 3
	SHI  Dice = 4
	GO   Dice = 5
	ROKU Dice = 6
)

func NewDices() *Dices {
	return &Dices{
		dice1: ICHI,
		dice2: ICHI,
	}
}

func (d *Dice) roll() {
	num := rand.Intn(6) + 1
	*d = Dice(num)
}

func (d *Dice) String() string {
	switch *d {
	case ICHI:
		return "ICHI"
	case NI:
		return "NI"
	case SAN:
		return "SAN"
	case SHI:
		return "SHI"
	case GO:
		return "GO"
	case ROKU:
		return "ROKU"
	default:
		panic("Unknown dice number")
	}
}

func (d *Dices) RollTheDice() {
	d.dice1.roll()
	d.dice2.roll()
}

func (d *Dices) Display(out io.Writer) {
	fmt.Fprintf(out, "    %4s - %4s\n", d.dice1.String(), d.dice2.String())
	fmt.Fprintf(out, "    %4d - %4d\n", d.dice1, d.dice2)
}

func (d *Dices) IsEven() bool {
	num1 := int(d.dice1)
	num2 := int(d.dice2)
	return (num1+num2)%2 == 0
}
