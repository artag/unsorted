package main

import (
	"fmt"
	"os"
	"time"

	i "dicemath/internal"

	"golang.org/x/term"
)

const (
	DiceWidth       = 9
	DiceHeight      = 5
	QuizDurationSec = 30
	MinDice         = 2
	MaxDice         = 6
	Reward          = 4
	Penalty         = 1
)

func main() {
	console := i.NewConsole(os.Stdin, os.Stdout)

	width, height, err := term.GetSize(0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	w := int(float64(width) * 0.6)
	h := int(float64(height) * 0.8)
	screen := i.NewScreen(w, h, os.Stdout)

	console.DisplayLine("Dice math")
	console.DisplayLine("")
	console.DisplayLine("You have %d seconds to answer as many as possible.", QuizDurationSec)
	console.DisplayLine("You get %d points for each correct answer.", Reward)
	console.DisplayLine("You lose %d points for each incorrect answer.", Penalty)
	console.DisplayLine("")
	console.DisplayLine("Press 'Enter' to begin...")
	console.PressEnter()

	runQuiz(console, screen)
}

func runQuiz(console *i.Console, screen *i.Screen) {
	correct := 0
	incorrect := 0

	for stay, timeout := true, time.After(QuizDurationSec*time.Second); stay; {
		if err := console.Clear(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		screen.Clear()
		dices := i.NewDices(screen, MinDice, MaxDice)
		sum := dices.GetSummaryValue()
		screen.Display()

		userInput := console.InputNumber()
		if sum == userInput {
			correct++
		} else {
			console.DisplayLine("Incorrect, the answer is %d", sum)
			time.Sleep(time.Second)
			incorrect++
		}

		select {
		case <-timeout:
			stay = false
		default:
		}
	}

	if err := console.Clear(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	score := correct*Reward - incorrect*Penalty
	console.DisplayLine("Correct:   %d", correct)
	console.DisplayLine("Incorrect: %d", incorrect)
	console.DisplayLine("Score:     %d", score)
}
