package main

import (
	"fmt"
	"io"
)

type Hand struct {
	out   io.Writer
	cards Cards
	name  string
}

func NewHand(out io.Writer, name string) *Hand {
	return &Hand{
		name:  name,
		cards: make(Cards, 0),
		out:   out}
}

func (h *Hand) Get(card Card) {
	h.cards = append(h.cards, card)
}

func (h *Hand) Put() *Card {
	if len(h.cards) < 1 {
		return nil
	}

	card := h.cards[0]
	h.cards = h.cards[1:]
	return &card
}

func (h *Hand) Display(hideFirstCard bool) {
	if hideFirstCard {
		fmt.Fprintf(h.out, "%s: ???\n", h.name)
	} else {
		value := h.cards.GetValue()
		fmt.Fprintf(h.out, "%s: %d\n", h.name, value)
	}

	for i := range h.cards {
		if hideFirstCard && i == 0 {
			h.cards[i].BackSide = true
		} else {
			h.cards[i].BackSide = false
		}
	}

	h.cards.Display(h.out)
	fmt.Fprintln(h.out)
}

func (h *Hand) CanDoubleDown(money int) bool {
	return len(h.cards) == 2 && money > 0
}

func (h *Hand) GetValue() int {
	return h.cards.GetValue()
}

func (h *Hand) Clear() {
	h.cards = make(Cards, 0)
}
