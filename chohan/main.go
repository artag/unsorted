package main

import (
	"bufio"
	"fmt"
	"os"

	c "chohan/classes"
)

const (
	InitialMoney     = 5000
	DealerFeePercent = 10
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = os.Stdout

	console = c.NewConsole(out)
	display = c.NewDisplay(out)
	input   = c.NewInput(in, out)
)

func main() {
	dealer := c.NewDealer(DealerFeePercent)
	dices := c.NewDices()
	player := c.NewPlayer(InitialMoney, false)

	display.WelcomeMessage()
	input.PressEnter("Press 'Enter' to continue...")

	if err := playGame(dealer, dices, player); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}

func playGame(dealer *c.Dealer, dices *c.Dices, player *c.Player) error {
	for {
		if err := console.Clear(); err != nil {
			return err
		}

		msg := fmt.Sprintf("You have %d mon. How much do you bet? (or (Q)UIT)", player.GetMoney())
		bet, quit := input.EnterBetOrQuit(msg, dealer, player)
		if quit {
			msg := fmt.Sprintf("You live the game with %d mon.", player.GetMoney())
			display.Message(msg)
			display.Message("Thanks for playing")
			os.Exit(0)
		}

		cho := input.EnterChoOrHan("CHO (even) or HAN (odd)?")
		display.Message("The dealer roll the dice...")
		dealer.RollTheDice(dices)
		dices.Display(out)

		isEven := dices.IsEven()
		if cho == isEven {
			playerWon(bet, dealer, player)
		} else {
			playerLoose(bet, player)
		}

		input.PressEnter("Press 'Enter to continue...'")
	}
}

func playerWon(bet int, dealer *c.Dealer, player *c.Player) {
	fmt.Printf("You won! You take %d mon.\n", bet)
	player.AddMoney(bet)
	fee := dealer.GetFee(bet)
	fmt.Printf("The house collects a %d mon fee.\n", fee)
	player.RemoveMoney(fee)
}

func playerLoose(bet int, player *c.Player) {
	fmt.Println("You lost!")
	player.RemoveMoney(bet)

	money := player.GetMoney()
	if money <= 0 {
		fmt.Println("You have run out of money!")
		fmt.Println("Thanks for playing!")
		os.Exit(0)
	}
}
