package main

import (
	"fmt"
	"io"
	"strconv"
)

var (
	Hearts   = fmt.Sprintf("%c", 9829)
	Diamonds = fmt.Sprintf("%c", 9830)
	Spades   = fmt.Sprintf("%c", 9824)
	Clubs    = fmt.Sprintf("%c", 9827)
	Backside = "backside"
)

type Card struct {
	Suit     string
	Rank     string
	BackSide bool
}

type Cards []Card

func (c *Card) Display(out io.Writer, row int) {
	if c.BackSide {
		switch row {
		case 0:
			fmt.Fprint(out, " ___  ")
		case 1:
			fmt.Fprint(out, "|## | ")
		case 2:
			fmt.Fprint(out, "|###| ")
		case 3:
			fmt.Fprint(out, "|_##| ")
		default:
			fmt.Fprintln(out, " ___ ")
			fmt.Fprintln(out, "|## |")
			fmt.Fprintln(out, "|###|")
			fmt.Fprintln(out, "|_##|")

		}
		return
	}

	switch row {
	case 0:
		fmt.Fprintf(out, " ___  ")
	case 1:
		fmt.Fprintf(out, "|%-2s | ", c.Rank)
	case 2:
		fmt.Fprintf(out, "| %s | ", c.Suit)
	case 3:
		if c.Rank == "10" {
			fmt.Fprintf(out, "|_%2s| ", c.Rank)
		} else {
			fmt.Fprintf(out, "|__%s| ", c.Rank)
		}
	default:
		fmt.Fprintln(out, " ___ ")
		fmt.Fprintf(out, "|%-2s |\n", c.Rank)
		fmt.Fprintf(out, "| %s |\n", c.Suit)
		if c.Rank == "10" {
			fmt.Fprintf(out, "|_%2s|\n", c.Rank)
		} else {
			fmt.Fprintf(out, "|__%s|\n", c.Rank)
		}
	}
}

func (c *Card) GetValue(sum int) int {
	if c.Rank == "A" {
		if sum+11 <= 21 {
			return 11
		}
		return 1
	} else if c.Rank == "K" || c.Rank == "Q" || c.Rank == "J" {
		return 10
	} else {
		num, _ := strconv.Atoi(c.Rank)
		return num
	}
}

func (c *Cards) Display(out io.Writer) {
	rows := []int{0, 1, 2, 3}
	for _, row := range rows {
		for _, card := range *c {
			card.Display(out, row)
		}
		fmt.Println()
	}
}

func (c *Cards) GetValue() int {
	sum := 0
	for _, card := range *c {
		sum += card.GetValue(sum)
	}

	return sum
}
