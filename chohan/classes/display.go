package classes

import (
	"fmt"
	"io"
)

type Display struct {
	out io.Writer
}

func NewDisplay(out io.Writer) *Display {
	return &Display{out: out}
}

func (d *Display) WelcomeMessage() {
	msg := fmt.Sprintf(
		"Cho-Han\n" +
			"In this traditional Japanese dice game, two dice are rolled in a bamboo\n" +
			"cup by the dealer sitting on the floor. The player must guess if the\n" +
			"dice total to an even (cho) or odd (han) number.\n")
	fmt.Fprintln(d.out, msg)
}

func (d *Display) Message(msg string) {
	fmt.Fprintln(d.out, msg)
}
