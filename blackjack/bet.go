package main

import (
	"fmt"
	"io"
)

type Bet struct {
	out   io.Writer
	Value int
}

func NewBet(out io.Writer, value int) *Bet {
	return &Bet{out: out, Value: value}
}

func (b *Bet) Add(additionalBet Bet) {
	b.Value += additionalBet.Value
}

func (b *Bet) Display() {
	fmt.Fprintf(b.out, "Bet: %d\n\n", b.Value)
}
