package main

import (
	"fmt"
	"math"
	"os"
	"time"

	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

const (
	NumberOfLogos   = 5
	MovingDelayInMs = 250
)

func main() {
	err := termbox.Init()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	width, height := termbox.Size()
	screenSize := NewScreenSize(width, height)
	logos := CreateLogos(screenSize, "DVD")

	bounces := 0
	go bgthread(logos, bounces)

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

func CreateLogos(size ScreenSize, text string) []*Logo {
	logos := make([]*Logo, NumberOfLogos)
	for i := 0; i < NumberOfLogos; i++ {
		logos[i] = NewLogo(size, text)
	}
	return logos
}

func tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x += runewidth.RuneWidth(c)
	}
}

func bgthread(logos []*Logo, bounces int) {
	ticker := time.NewTicker(MovingDelayInMs * time.Millisecond)
	defer ticker.Stop()
	for {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		screenSize := NewScreenSize(termbox.Size())

		for _, logo := range logos {
			tbprint(logo.X, logo.Y, logo.color, termbox.ColorDefault, logo.text)
			logo.Move(screenSize, &bounces)
		}

		bounces %= math.MaxInt
		msg := fmt.Sprintf("Corner bounces: %d", bounces)
		tbprint(1, 1, termbox.ColorDefault, termbox.ColorDefault, msg)
		tbprint(1, 2, termbox.ColorDefault, termbox.ColorDefault, "Press 'Esc' or 'Ctrl+C' to quit...")

		termbox.Flush()
		<-ticker.C // wait for ticker
	}
}
