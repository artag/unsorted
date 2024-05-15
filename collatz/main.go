package main

import (
	"bufio"
	"os"

	c "collatz/classes"
)

const (
	// Перенос на следующую строку после n-го слова
	HyphenateAfterCount = 16
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = os.Stdout
)

func main() {
	input := c.NewInput(in, out)
	display := c.NewDisplay(out, HyphenateAfterCount)

	number, quit := input.EnterPositiveNumberOrQuit(
		"Enter a starting number (greater than 0) or (Q)UIT:")
	if quit {
		os.Exit(0)
	}

	cnt := 0
	display.Number(number, cnt)

	for number != 1 {
		if number%2 == 0 {
			number = number / 2
		} else {
			number = 3*number + 1
		}

		cnt++
		display.Number(number, cnt)
	}
}
