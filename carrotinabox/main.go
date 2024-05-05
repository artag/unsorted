package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"

	c "carrotinabox/classes"
)

var (
	out = os.Stdout
	in  = bufio.NewReader(os.Stdin)
)

func main() {
	input := c.NewInput(out, in)
	display := c.NewDisplay(out)
	console := c.NewConsole(out)

	display.WelcomeMessage()
	input.PressEnter("Press 'Enter' to begin...")
	if err := console.Clear(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var player1 *c.Player
	var player2 *c.Player
	human1 := input.EnterHumanOrComputer("Player 1 is (H)uman or (C)omputer?")
	if human1 {
		p1Name := input.EnterName("Human player 1, enter your name: ")
		player1 = c.NewPlayer(p1Name, false)
	} else {
		p1Name := "Computer1"
		player1 = c.NewPlayer(p1Name, true)
	}

	human2 := input.EnterHumanOrComputer("Player 2 is (H)uman or (C)omputer?")
	if human2 {
		p2Name := input.EnterName("Human player 2, enter your name: ")
		player2 = c.NewPlayer(p2Name, false)
	} else {
		p2Name := "Computer2"
		player2 = c.NewPlayer(p2Name, true)
	}

	for {
		console.Clear()
		display.Scores(*player1, *player2)
		display.TwoClosedBoxes()
		display.PlayerNames(*player1, *player2)

		display.GameMessage(*player1, *player2)
		input.PressEnter("When " + player2.Name() + " has closed their eyes, press Enter...")
		console.Clear()

		carrotInTheFirstBox := rand.Intn(2) == 0

		if player1.IsHuman() {
			fmt.Println(player1.Name() + " here is the inside of your box:")
			display.TwoBoxesWithOneOpen(carrotInTheFirstBox)
			input.PressEnter("Press 'Enter' to continue...")
			console.Clear()

			fmt.Println(player1.Name() + " tell " + player2.Name() + " to open their eyes.")
			input.PressEnter("Press 'Enter' to continue...")

			fmt.Println(player1.Name() + ", say one of the following sentences to " + player2.Name() + ":")
			fmt.Println("    1) There is a carrot in my box.")
			fmt.Println("    2) There is not a carrot in my box.")
			input.PressEnter("Press 'Enter' to continue...")
		} else {
			fmt.Println("Thinking...")
			sleep()
			sentence := rand.Intn(2)
			if sentence == 0 {
				fmt.Println("There is a carrot in my box.")
			} else {
				fmt.Println("There is not a carrot in my box.")
			}
		}

		firstBox := "RED "
		secondBox := "GOLD"
		var answerYes bool
		if player2.IsHuman() {
			answerYes = input.EnterYesNo(
				player2.Name() + ", do you want to swap boxes with " + player1.Name() + "? (Y)es/(N)o")
		} else {
			fmt.Println("Swap boxes or not... Thinking...")
			sleep()
			answerYes = rand.Intn(2) == 0
			if answerYes {
				fmt.Println("Swapping!")
			} else {
				fmt.Println("Leave unchanged!")
			}
		}

		if answerYes {
			carrotInTheFirstBox = !carrotInTheFirstBox
			firstBox, secondBox = secondBox, firstBox
		}
		display.TwoOpenBoxes(firstBox, secondBox, carrotInTheFirstBox)

		if carrotInTheFirstBox {
			fmt.Println(player1.Name() + " is the winner!")
			player1.AddScore()
		} else {
			fmt.Println(player2.Name() + " is the winner!")
			player2.AddScore()
		}

		playAgain := input.EnterYesNo("Do you want to play again? (Y)es/(N)o")
		if !playAgain {
			break
		}
	}

	fmt.Println()
	display.Scores(*player1, *player2)
	fmt.Println("Thanks for playing!")
}

func sleep() {
	seconds := rand.Intn(10)
	seconds++
	time.Sleep(time.Duration(seconds) * time.Second)
}
