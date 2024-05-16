package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

type Screen struct {
	Width  int
	Height int
}

const (
	Width           = 79
	Height          = 20
	DEAD            = ' '
	MovingDelayInMs = 100
)

func main() {
	redefineFlagUsage()
	seedCount := flag.Int("seed", 1, "Seed count. 1 maximum seed count. 2 or more - lower seed count.")
	symbolCode := flag.Int("symbol", 79, "Cells symbol ASCII code. Maybe you want to try codes: 9619, 35.")
	flag.Parse()

	if flag.NFlag() < 1 {
		flag.Usage()
		pressEnterToContinue()
	}

	aliveSymbol := rune(*symbolCode)

	err := termbox.Init()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	width, height := termbox.Size()
	screen := Screen{Width: width, Height: height}

	cells := initCells(screen)
	nextCells := initCells(screen)
	seedCells(nextCells, screen, *seedCount, aliveSymbol)

	go bgthread(cells, nextCells, screen, aliveSymbol)

	for {
		ev := termbox.PollEvent()
		if ev.Type == termbox.EventKey {
			if ev.Key == termbox.KeyEsc || ev.Key == termbox.KeyCtrlC {
				break
			}

			termbox.Flush()
		}
	}

	termbox.Close()
}

func redefineFlagUsage() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Conway's Game of Life\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Usage information:\n")
		flag.PrintDefaults()
	}
}

func pressEnterToContinue() {
	fmt.Println("Press 'Enter' to continue...")
	in := bufio.NewReader(os.Stdin)
	_, _ = in.ReadString('\n')
	fmt.Println()
}

func initCells(screen Screen) [][]rune {
	cells := make([][]rune, screen.Width)
	for i := range cells {
		cells[i] = make([]rune, screen.Width)
	}
	return cells
}

func seedCells(nextCells [][]rune, screen Screen, seedCount int, aliveSymbol rune) {
	var seed int
	if seedCount < 1 {
		seed = 2
	} else {
		seed = seedCount + 1
	}

	for x := 0; x < screen.Width; x++ {
		for y := 0; y < screen.Height; y++ {
			live := rand.Intn(seed) == 0
			if live {
				nextCells[x][y] = aliveSymbol
			} else {
				nextCells[x][y] = DEAD
			}
		}
	}
}

func copyCells(cells, nextCells [][]rune, screen Screen) {
	for x := 0; x < screen.Width; x++ {
		for y := 0; y < screen.Height; y++ {
			cells[x][y] = nextCells[x][y]
		}
	}
}

func bgthread(cells, nextCells [][]rune, screen Screen, alive rune) {
	ticker := time.NewTicker(MovingDelayInMs * time.Millisecond)
	defer ticker.Stop()
	for {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		copyCells(cells, nextCells, screen)

		for x := 0; x < screen.Width; x++ {
			for y := 0; y < screen.Height; y++ {
				tbprintChar(x, y, termbox.ColorWhite, termbox.ColorDefault, cells[x][y])
			}
		}

		for x := 0; x < screen.Width; x++ {
			for y := 0; y < screen.Height; y++ {
				numNeighbors := countCellNeighbors(cells, x, y, screen, alive)

				if cells[x][y] == alive && (numNeighbors == 2 || numNeighbors == 3) {
					nextCells[x][y] = alive
				} else if cells[x][y] == DEAD && numNeighbors == 3 {
					nextCells[x][y] = alive
				} else {
					nextCells[x][y] = DEAD
				}
			}
		}

		tbprint(1, screen.Height-1, termbox.ColorDefault, termbox.ColorDefault, "Press 'Esc' or 'Ctrl+C' to quit...")

		termbox.Flush()
		<-ticker.C // wait for ticker
	}
}

func countCellNeighbors(cells [][]rune, x, y int, screen Screen, alive rune) int {
	var left int
	if x-1 < 0 {
		left = screen.Width - 1
	} else {
		left = x - 1
	}

	var right int
	if x+1 >= screen.Width {
		right = 0
	} else {
		right = x + 1
	}

	var above int
	if y-1 < 0 {
		above = screen.Height - 1
	} else {
		above = y - 1
	}

	var below int
	if y+1 >= screen.Height {
		below = 0
	} else {
		below = y + 1
	}

	numNeighbors := 0
	if cells[left][above] == alive {
		numNeighbors++
	}
	if cells[x][above] == alive {
		numNeighbors++
	}
	if cells[right][above] == alive {
		numNeighbors++
	}
	if cells[left][y] == alive {
		numNeighbors++
	}
	if cells[right][y] == alive {
		numNeighbors++
	}
	if cells[left][below] == alive {
		numNeighbors++
	}
	if cells[x][below] == alive {
		numNeighbors++
	}
	if cells[right][below] == alive {
		numNeighbors++
	}

	return numNeighbors
}

func tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x += runewidth.RuneWidth(c)
	}
}

func tbprintChar(x, y int, fg, bg termbox.Attribute, char rune) {
	termbox.SetCell(x, y, char, fg, bg)
}
