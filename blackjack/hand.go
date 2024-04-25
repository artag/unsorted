package main

import (
	"fmt"
	"io"
)

type Hand struct {
	cards Cards
	name  string
}

func NewHand(name string) *Hand {
	return &Hand{name: name, cards: make(Cards, 0)}
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

func (h *Hand) Display(out io.Writer, hideFirstCard bool) {
	if hideFirstCard {
		fmt.Fprintf(out, "%s: ???\n", h.name)
	} else {
		value := h.cards.GetValue()
		fmt.Fprintf(out, "%s: %d\n", h.name, value)
	}

	for i := range h.cards {
		if hideFirstCard && i == 0 {
			h.cards[i].BackSide = true
		} else {
			h.cards[i].BackSide = false
		}
	}

	h.cards.Display(out)
	fmt.Fprintln(out)
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
