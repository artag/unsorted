package main

import (
	"math/rand"

	"github.com/nsf/termbox-go"
)

var directionX = []int{-1, 1}
var directionY = []int{-1, 1}

var speedX = []int{1, 2, 3}
var speedY = []int{1, 2, 3}

var colors = []termbox.Attribute{
	termbox.ColorBlue,
	termbox.ColorCyan,
	termbox.ColorDarkGray,
	termbox.ColorGreen,
	termbox.ColorLightBlue,
	termbox.ColorLightCyan,
	termbox.ColorLightGray,
	termbox.ColorLightGreen,
	termbox.ColorLightMagenta,
	termbox.ColorLightRed,
	termbox.ColorLightYellow,
	termbox.ColorMagenta,
	termbox.ColorRed,
	termbox.ColorWhite,
	termbox.ColorYellow,
}

type Logo struct {
	X          int
	Y          int
	speedX     int
	speedY     int
	directionX int
	directionY int
	color      termbox.Attribute
	text       string
	logoLength int
	logoHeight int
}

type ScreenSize struct {
	Width  int
	Height int
}

func NewScreenSize(width, height int) ScreenSize {
	return ScreenSize{Width: width, Height: height}
}

func NewLogo(screenSize ScreenSize, text string) *Logo {
	textLen := len(text)
	return &Logo{
		X:          randRange(1, screenSize.Width),
		Y:          randRange(1, screenSize.Height),
		speedX:     randItem(speedX),
		speedY:     randItem(speedY),
		directionX: randItem(directionX),
		directionY: randItem(directionY),
		color:      randItem(colors),
		text:       text,
		logoLength: textLen,
		logoHeight: 1,
	}
}

func (l *Logo) Move(size ScreenSize, bounces *int) {
	if l.X >= size.Width-l.logoLength {
		l.directionX = -1
		l.color = randItem(colors)
		*bounces += 1
	}
	if l.X <= 1 {
		l.directionX = 1
		l.color = randItem(colors)
		*bounces += 1
	}
	if l.Y >= size.Height-l.logoHeight {
		l.directionY = -1
		l.color = randItem(colors)
		*bounces += 1
	}
	if l.Y <= 1 {
		l.directionY = 1
		l.color = randItem(colors)
		*bounces += 1
	}

	l.X += l.speedX * l.directionX
	l.Y += l.speedY * l.directionY
}

func randRange(min, max int) int {
	return rand.Intn(max-min) + min
}

func randItem[T any](arr []T) T {
	len := len(arr)
	idx := rand.Intn(len)
	return arr[idx]
}
