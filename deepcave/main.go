package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

const (
	MovingDelayInMs = 50
)

func main() {
	fmt.Println("Deep cave")
	pressEnter()

	err := termbox.Init()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	width, height := termbox.Size()
	screen := NewScreen(width, height)

	go bgthread(screen)

	for {
		ev := termbox.PollEvent()
		if ev.Type == termbox.EventKey {
			if ev.Key == termbox.KeyCtrlC {
				break
			}
			termbox.Flush()
		}
	}

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.Flush()
	termbox.Close()
}

func bgthread(screen *Screen) {
	gapWidth := 10
	leftWidth := (screen.Width - gapWidth) / 2

	// Initial rows
	for i := 0; i < screen.Height-2; i++ {
		row := createRow(&gapWidth, &leftWidth, screen.Width)
		screen.Push(row)
	}

	ticker := time.NewTicker(MovingDelayInMs * time.Millisecond)
	defer ticker.Stop()

	// Main cycle
	for {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

		// Display
		screen.Display(tbprintChar)
		tbprintString(0, screen.Height-1, "Press 'Ctrl-C' to quit.")
		termbox.Flush()

		// Wait for ticker
		<-ticker.C

		// Delete upper row, create and add bottom row
		screen.Pop()
		bottomRow := createRow(&gapWidth, &leftWidth, screen.Width)
		screen.Push(bottomRow)
	}
}

func createRow(gapWidth *int, leftWidth *int, width int) string {
	// Build row
	rightWidth := width - *gapWidth - *leftWidth
	var sb strings.Builder
	for i := 0; i < *leftWidth; i++ {
		sb.WriteRune('#')
	}
	for i := 0; i < *gapWidth; i++ {
		sb.WriteRune(' ')
	}
	for i := 0; i < rightWidth; i++ {
		sb.WriteRune('#')
	}

	// Randomly change left width
	diceRoll := rand.Intn(6)
	if diceRoll == 1 && *leftWidth > 1 {
		*leftWidth--
	} else if diceRoll == 2 && *leftWidth+*gapWidth < width-1 {
		*leftWidth++
	}

	// Randomly change gap width
	diceRoll2 := rand.Intn(6)
	if diceRoll2 == 1 && *gapWidth > 2 {
		*gapWidth--
	} else if diceRoll2 == 2 && *leftWidth+*gapWidth < width-1 {
		*gapWidth++
	}

	// Convert row to string
	return sb.String()
}

func pressEnter() {
	in := bufio.NewReader(os.Stdin)
	fmt.Println("Press 'Enter' to start.")
	_, _ = in.ReadString('\n')
	fmt.Println()
}

func tbprintString(x, y int, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, termbox.ColorDefault, termbox.ColorDefault)
		x += runewidth.RuneWidth(c)
	}
}

func tbprintChar(x, y int, char rune) {
	termbox.SetCell(x, y, char, termbox.ColorDefault, termbox.ColorDefault)
}
