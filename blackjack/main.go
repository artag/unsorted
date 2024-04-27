package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var (
	InitialMoney = 5000
	MinBet       = 1
	MaxValue     = 21
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = os.Stdout
)

func main() {

	console := NewConsole(out)

	money := NewMoney(out, InitialMoney)
	dealerHand := NewHand(out, "DEALER")
	playerHand := NewHand(out, "PLAYER")

	displayWelcomeMessage(out)

	for {
		// Clear console
		if err := console.Clear(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		PlayRound(playerHand, dealerHand, money)
	}
}

func PlayRound(playerHand, dealerHand *Hand, money *Money) {
	if money.IsNotEnough() {
		fmt.Fprintln(out,
			"You're broke!\n"+
				"Good thing you weren't playing with real money.\n"+
				"Thanks for playing!")
		os.Exit(0)
	}

	playerHand.Clear()
	dealerHand.Clear()
	deck := NewDeck()
	deck.Shuffle()

	money.Display()
	quit, bet := getPlayerBet(money.Value)
	if quit {
		showMessageAndQuit()
	}
	bet.Display()

	dealerHand.Get(*deck.Pop())
	dealerHand.Get(*deck.Pop())
	playerHand.Get(*deck.Pop())
	playerHand.Get(*deck.Pop())

	gamePlayer(deck, playerHand, dealerHand, money, bet)
	gameDealer(deck, playerHand, dealerHand)
	endRound(playerHand, dealerHand, bet, money)
}

func gamePlayer(deck *Deck, playerHand, dealerHand *Hand, money *Money, bet *Bet) {
	for {
		displayHands(*playerHand, *dealerHand, true)

		if playerHand.GetValue() > MaxValue {
			break
		}

		move := getPlayerMove(*playerHand, money.Value-bet.Value)

		if move == DoubleDown {
			remains := min(bet.Value, (money.Value - bet.Value))
			quit, additionalBet := getPlayerBet(remains)
			if quit {
				showMessageAndQuit()
			}
			bet.Add(*additionalBet)
			bet.Display()
		}

		if move == DoubleDown || move == Hit {
			newCard := deck.Pop()
			playerHand.Get(*newCard)
			if playerHand.GetValue() > MaxValue {
				continue
			}
		}

		if move == DoubleDown || move == Stand {
			break
		}
	}
}

func gameDealer(deck *Deck, playerHand *Hand, dealerHand *Hand) {
	if playerHand.GetValue() > MaxValue {
		return
	}

	for dealerHand.GetValue() < 17 {
		fmt.Fprintf(out, "Dealer hits...\n\n")

		card := *deck.Pop()
		dealerHand.Get(card)

		displayHands(*playerHand, *dealerHand, true)
		if dealerHand.GetValue() > MaxValue {
			break
		}

		waitToPressEnter("Press Enter to continue...")
	}
}

func endRound(playerHand, dealerHand *Hand, bet *Bet, money *Money) {
	displayHands(*playerHand, *dealerHand, false)
	playerValue := playerHand.GetValue()
	dealerValue := dealerHand.GetValue()

	if dealerValue > MaxValue {
		fmt.Fprintf(out, "Dealer busts! You win %d!\n", bet.Value)
		money.Add(*bet)
	} else if playerValue > MaxValue || playerValue < dealerValue {
		fmt.Fprintln(out, "You lost!")
		money.Take(*bet)
	} else if playerValue > dealerValue {
		fmt.Fprintf(out, "You won %d!\n", bet.Value)
		money.Add(*bet)
	} else {
		fmt.Fprintln(out, "It's a tie, the bet is returned to you.")
	}

	waitToPressEnter("Press Enter to continue...")
}

func displayWelcomeMessage(out io.Writer) {
	msg := "Blackjack.\n" +
		"Rules:\n" +
		"    Try to get as close to 21 without going over.\n" +
		"    Kings, Queens, and Jacks are worth 10 points.\n" +
		"    Aces are worth 1 or 11 points.\n" +
		"    Cards 2 through 10 are worth their face value.\n" +
		"    (H)it to take another card.\n" +
		"    (S)tand to stop taking cards.\n" +
		"    On your first play, you can (D)ouble down to increase your bet\n" +
		"    but must hit exactly one more time before standing.\n" +
		"    In case of a tie, the bet is returned to the player.\n" +
		"    The dealer stops hitting at 17.\n"
	fmt.Fprintln(out, msg)
}

func displayHands(playerHand, dealerHand Hand, hideDealerHand bool) {
	dealerHand.Display(hideDealerHand)
	playerHand.Display(false)
	fmt.Fprintln(out)
}

func getPlayerBet(maxBet int) (bool, *Bet) {
	for {
		fmt.Fprintf(out, "How much do you bet? (%d-%d, or QUIT)\n", MinBet, maxBet)
		fmt.Fprint(out, "> ")
		input, err := in.ReadString('\n')
		input = strings.TrimRight(input, "\r\n")
		if err != nil {
			continue
		}

		betStr := strings.TrimSuffix(input, "\n")
		str := strings.ToLower(betStr)
		if str == "quit" {
			return true, NewBet(out, 0)
		}

		bet, err := strconv.Atoi(betStr)
		if err != nil {
			continue
		}

		if MinBet < bet && bet <= maxBet {
			fmt.Fprintln(out)
			return false, NewBet(out, bet)
		}
	}
}

func getPlayerMove(hand Hand, money int) Move {
	for {
		canDoubleDown := hand.CanDoubleDown(money)
		if canDoubleDown {
			fmt.Fprintln(out, "(H)it, (S)tand, (D)ouble down")
		} else {
			fmt.Fprintln(out, "(H)it, (S)tand")
		}
		fmt.Fprint(out, "> ")

		input, err := in.ReadString('\n')
		input = strings.TrimRight(input, "\r\n")
		if err != nil {
			continue
		}

		inputStr := strings.TrimSuffix(input, "\n")
		str := strings.ToLower(inputStr)
		if canDoubleDown && str == "d" {
			fmt.Fprintln(out)
			return DoubleDown
		}
		if str == "s" {
			fmt.Fprintln(out)
			return Stand
		}
		if str == "h" {
			fmt.Fprintln(out)
			return Hit
		}
	}
}

func waitToPressEnter(message string) {
	fmt.Fprint(out, message)
	_, _ = in.ReadString('\n')
	fmt.Fprintln(out)
}

func showMessageAndQuit() {
	fmt.Fprintln(out, "Thanks for playing!")
	os.Exit(0)
}
