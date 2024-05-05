package classes

import (
	"fmt"
	"io"
	"strings"
)

type Display struct {
	out io.Writer
}

func NewDisplay(out io.Writer) *Display {
	return &Display{
		out: out,
	}
}

func (d *Display) WelcomeMessage() {
	fmt.Fprintln(d.out, "Carrot in a Box.\n\n"+
		"This is a bluffing game for two human players. Each player has a box.\n"+
		"One box has a carrot in it. To win, you must have the box with the carrot in it.\n\n"+
		"This is a very simple and silly game.\n\n"+
		"The first player looks into their box (the second player must close\n"+
		"their eyes during this). The first player then says \"There is a carrot\n"+
		"in my box\" or \"There is not a carrot in my box\". The second player then\n"+
		"gets to decide if they want to swap boxes or not.")
}

func (d *Display) TwoClosedBoxes() {
	boxes := `HERE ARE TWO BOXES:
  __________     __________
 /         /|   /         /|
+---------+ |  +---------+ |
|   RED   | |  |   GOLD  | |
|   BOX   | /  |   BOX   | /
+---------+/   +---------+/`
	fmt.Println(boxes)
}

func (d *Display) TwoBoxesWithOneOpen(carrotInFirstBox bool) {
	var boxes string
	if carrotInFirstBox {
		boxes = `
   ___VV____
  |   VV    |
  |   VV    |
  |___||____|    __________
 /    ||   /|   /         /|
+---------+ |  +---------+ |
|   RED   | |  |   GOLD  | |
|   BOX   | /  |   BOX   | /
+---------+/   +---------+/
 (carrot!)`
	} else {
		boxes = `
   _________
  |         |
  |         |
  |_________|    __________
 /         /|   /         /|
+---------+ |  +---------+ |
|   RED   | |  |   GOLD  | |
|   BOX   | /  |   BOX   | /
+---------+/   +---------+/
 (no carrot!)`
	}
	fmt.Fprintln(d.out, boxes)
}

func (d *Display) TwoOpenBoxes(firstBox, secondBox string, carrotInFirstBox bool) {
	var out strings.Builder
	if carrotInFirstBox {
		out.WriteString(fmt.Sprintln("   ___VV____      _________ "))
		out.WriteString(fmt.Sprintln("  |   VV    |    |         |"))
		out.WriteString(fmt.Sprintln("  |   VV    |    |         |"))
		out.WriteString(fmt.Sprintln("  |___||____|    |_________|"))
		out.WriteString(fmt.Sprintln(" /    ||   /|   /         /|"))
		out.WriteString(fmt.Sprintln("+---------+ |  +---------+ |"))
		out.WriteString(fmt.Sprintf("|   %-4s  | |  |   %-4s  | |\n", firstBox, secondBox))
		out.WriteString(fmt.Sprintln("|   BOX   | /  |   BOX   | /"))
		out.WriteString(fmt.Sprintln("+---------+/   +---------+/"))
	} else {
		out.WriteString(fmt.Sprintln("   _________      ___VV____ "))
		out.WriteString(fmt.Sprintln("  |         |    |   VV    |"))
		out.WriteString(fmt.Sprintln("  |         |    |   VV    |"))
		out.WriteString(fmt.Sprintln("  |_________|    |___||____|"))
		out.WriteString(fmt.Sprintln(" /         /|   /    ||   /|"))
		out.WriteString(fmt.Sprintln("+---------+ |  +---------+ |"))
		out.WriteString(fmt.Sprintf("|   %-4s  | |  |   %-4s  | |\n", firstBox, secondBox))
		out.WriteString(fmt.Sprintln("|   BOX   | /  |   BOX   | /"))
		out.WriteString(fmt.Sprintln("+---------+/   +---------+/"))
	}

	fmt.Fprintln(d.out, out.String())
}

func (d *Display) GameMessage(player1, player2 Player) {
	fmt.Fprintln(d.out, player1.Name()+", you have a RED box in front of you.")
	fmt.Fprintln(d.out, player2.Name()+", you have a GOLD box in front of you.\n")
	fmt.Fprintln(d.out, player1.Name()+", you will get to look into your box.")
	fmt.Fprintf(d.out, strings.ToUpper(player2.Name())+", close your eyes and don't look!!!\n")
}

func (d *Display) Scores(player1, player2 Player) {
	fmt.Fprintf(
		d.out,
		"Game scores.\n%s\n%s\n\n",
		player1.NameWithScore(), player2.NameWithScore())
}

func (d *Display) PlayerNames(player1, player2 Player) {
	const boxWidth = 11
	const boxSpace = 4

	len1 := getStringLenth(player1.Name())
	len2 := getStringLenth(player2.Name())
	indent1Total := absInt((boxWidth - len1) / 2)
	indent2 := absInt((boxWidth - len2) / 2)
	indent2Total := boxWidth - (indent1Total + len1) + boxSpace + indent2

	var out strings.Builder
	for i := 0; i < indent1Total; i++ {
		out.WriteRune(' ')
	}
	out.WriteString(player1.Name())
	for i := 0; i < indent2Total; i++ {
		out.WriteRune(' ')
	}
	out.WriteString(player2.Name())
	out.WriteRune('\n')
	fmt.Fprintln(d.out, out.String())
}

func getStringLenth(str string) int {
	split := strings.Split(str, "")
	return len(split)
}

func absInt(x int) int {
	return absDiffInt(x, 0)
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
