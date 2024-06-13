package internal

import "math/rand"

type Dices []*Dice

func NewDices(screen *Screen, minNumberOfDices, maxNumberOfDices int) Dices {
	width := screen.Width
	height := screen.Height

	dices := make([]*Dice, 0)
	dice := NewDice(width, height)
	dices = append(dices, dice)
	screen.Add(dice)

	num := 1
	maxNum := rand.Intn(maxNumberOfDices-1) + minNumberOfDices
	for num < maxNum {
		addNewDice := true
		newDice := NewDice(width, height)
		for _, availDice := range dices {
			if availDice.IsOverlaps(newDice) {
				addNewDice = false
			}
		}

		if addNewDice {
			dices = append(dices, newDice)
			screen.Add(newDice)
			num++
		}
	}

	return Dices(dices)
}

func (d *Dices) GetSummaryValue() int {
	sum := 0
	for _, dice := range *d {
		sum += dice.diceType.value
	}

	return sum
}
