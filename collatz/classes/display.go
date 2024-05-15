package classes

import (
	"fmt"
	"io"
)

type Display struct {
	out       io.Writer
	hyphenate int
}

func NewDisplay(out io.Writer, hyphenate int) *Display {
	var h int
	if hyphenate < 4 {
		h = 4
	} else {
		h = hyphenate
	}

	return &Display{
		out:       out,
		hyphenate: h,
	}
}

func (d *Display) WelcomeMessage() {
	msg := fmt.Sprintf(
		"Collatz Sequence, or, the 3n + 1 Problem\n\n" +
			"The Collatz sequence is a sequence of numbers produced from a starting\n" +
			"number n, following three rules:\n" +
			"1) If n is even, the next number n is n / 2.\n" +
			"2) If n is odd, the next number n is n * 3 + 1.\n" +
			"3) If n is 1, stop. Otherwise, repeat.\n\n" +
			"It is generally thought, but so far not mathematically proven, that\n" +
			"every starting number eventually terminates at 1.\n")
	fmt.Fprintln(d.out, msg)
}

func (d *Display) Number(num int, cnt int) {
	// Перенос на другую строку
	if cnt != 0 && cnt%d.hyphenate == 0 {
		fmt.Fprintln(d.out)
	}

	if num == 1 {
		fmt.Fprintf(d.out, "%d\n", num)
	} else {
		fmt.Fprintf(d.out, "%d, ", num)
	}
}
