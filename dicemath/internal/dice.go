package internal

import (
	"math/rand"
)

const (
	diceMaxValue = 6
	diceWidth    = 9
	diceHeight   = 5
)

type Dice struct {
	topLeft     Coord
	topRight    Coord
	bottomLeft  Coord
	bottomRight Coord
	diceType    diceType
}

func NewDice(xmax, ymax int) *Dice {
	x := rand.Intn(xmax - diceWidth - 1)
	y := rand.Intn(ymax - diceHeight - 1)
	diceTypeIndex := rand.Intn(len(allDiceTypes))
	t := allDiceTypes[diceTypeIndex]
	return &Dice{
		topLeft:     Coord{x: x, y: y},
		topRight:    Coord{x: x + diceWidth, y: y},
		bottomLeft:  Coord{x: x, y: y + diceHeight},
		bottomRight: Coord{x: x + diceWidth, y: y + diceHeight},
		diceType:    t,
	}
}

func (d1 *Dice) IsOverlaps(d2 *Dice) bool {
	if d1.topLeft.x <= d2.topLeft.x &&
		d2.topLeft.x <= d1.topRight.x &&
		d1.topLeft.y <= d2.topLeft.y &&
		d2.topLeft.y <= d1.bottomLeft.y {
		return true
	}

	if d1.topLeft.x <= d2.topRight.x &&
		d2.topRight.x <= d1.topRight.x &&
		d1.topLeft.y <= d2.topRight.y &&
		d2.topRight.y <= d1.bottomLeft.y {
		return true
	}

	if d1.topLeft.x <= d2.bottomLeft.x &&
		d2.bottomLeft.x <= d1.topRight.x &&
		d1.topLeft.y <= d2.bottomLeft.y &&
		d2.bottomLeft.y <= d1.bottomLeft.y {
		return true
	}

	if d1.topLeft.x <= d2.bottomRight.x &&
		d2.bottomRight.x <= d1.topRight.x &&
		d1.topLeft.y <= d2.bottomRight.y &&
		d2.bottomRight.y <= d1.bottomLeft.y {
		return true
	}

	return false
}

func (d *Dice) GetChars() map[Coord]rune {
	rows := d.diceType.rows
	m := make(map[Coord]rune)
	for y, row := range rows {
		for x, char := range row {
			coord := Coord{x: x + d.topLeft.x, y: y + d.topLeft.y}
			m[coord] = char
		}
	}

	return m
}
