package internal

import (
	"fmt"
	"io"
	"strings"
)

type Printer func(msg string)

type Coord struct {
	x int
	y int
}

type Screen struct {
	chars  map[Coord]rune
	dices  []*Dice
	out    io.Writer
	Width  int
	Height int
}

func NewScreen(width, height int, out io.Writer) *Screen {
	chars := clearScreen(height, width)
	return &Screen{
		chars:  chars,
		dices:  make([]*Dice, 0),
		out:    out,
		Width:  width,
		Height: height,
	}
}

func (s *Screen) Add(dice *Dice) {
	s.dices = append(s.dices, dice)
	diceChars := dice.GetChars()
	for diceCoord, diceChar := range diceChars {
		s.chars[diceCoord] = diceChar
	}
}

func (s *Screen) Clear() {
	s.chars = clearScreen(s.Height, s.Width)
	s.dices = make([]*Dice, 0)
}

func (s *Screen) Display() {
	var row strings.Builder
	for y := 0; y < s.Height; y++ {
		row.Reset()
		for x := 0; x < s.Width; x++ {
			coord := Coord{x: x, y: y}
			char := s.chars[coord]
			row.WriteRune(char)
		}
		str := row.String()
		fmt.Fprintln(s.out, str)
	}
}

func clearScreen(height int, width int) map[Coord]rune {
	chars := make(map[Coord]rune)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			coord := Coord{x: x, y: y}
			chars[coord] = ' '
		}
	}
	return chars
}
