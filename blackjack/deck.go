package main

import (
	"math/rand"
)

type Deck struct {
	cards []Card
}

func NewDeck() *Deck {
	deck := Deck{}
	for _, suit := range []string{Hearts, Diamonds, Spades, Clubs} {
		for _, rank := range []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"} {
			card := Card{Suit: suit, Rank: rank, BackSide: true}
			deck.Push(card)
		}
	}

	return &deck
}

func (d *Deck) Push(card Card) {
	d.cards = append(d.cards, card)
}

func (d *Deck) Pop() *Card {
	if len(d.cards) < 1 {
		return nil
	}

	card := d.cards[0]
	d.cards = d.cards[1:]
	return &card
}

func (d *Deck) Shuffle() {
	cnt := len(d.cards)
	for i := range d.cards {
		j := rand.Intn(cnt)
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	}
}
