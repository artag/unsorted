package main

import (
	"fmt"
	"io"
)

type Money struct {
	out   io.Writer
	Value int
}

func NewMoney(out io.Writer, value int) *Money {
	return &Money{out: out, Value: value}
}

func (m *Money) Add(bet Bet) {
	m.Value += bet.Value
}

func (m *Money) Take(bet Bet) {
	m.Value -= bet.Value
}

func (m *Money) IsEnough() bool {
	return m.Value > 0
}

func (m *Money) IsNotEnough() bool {
	return !m.IsEnough()
}

func (m *Money) Display() {
	fmt.Fprintf(m.out, "Money: %d\n", m.Value)
}
