package classes

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Input struct {
	in  *bufio.Reader
	out io.Writer
}

func NewInput(in *bufio.Reader, out io.Writer) *Input {
	return &Input{
		in:  in,
		out: out,
	}
}

func (i *Input) PressEnter(msg string) {
	fmt.Fprintln(i.out, msg)
	_, _ = i.in.ReadString('\n')
	fmt.Fprintln(i.out)
}

func (i *Input) EnterBetOrQuit(msg string, dealer *Dealer, player *Player) (int, bool) {
	for {
		fmt.Fprintln(i.out, msg)
		fmt.Fprint(i.out, "> ")

		input, err := i.in.ReadString('\n')
		input = strings.TrimRight(input, "\r\n")
		if err != nil {
			continue
		}
		if len(input) < 1 {
			continue
		}

		inputStr := strings.ToLower(input)
		if strings.HasPrefix(inputStr, "q") {
			return -1, true
		}

		bet, err := strconv.Atoi(input)
		if err != nil {
			continue
		}
		canBet := player.CanBet(bet)
		if !canBet {
			fmt.Fprintln(i.out, "You don't have enough to make that bet.")
			continue
		}

		return bet, false
	}
}

func (i *Input) EnterChoOrHan(msg string) bool {
	for {
		fmt.Fprintln(i.out, msg)
		fmt.Fprint(i.out, "> ")

		input, err := i.in.ReadString('\n')
		input = strings.TrimRight(input, "\r\n")
		if err != nil {
			continue
		}
		if len(input) < 1 {
			fmt.Fprintln(i.out, "Please enter either 'CHO' or 'HAN'.")
			continue
		}

		inputStr := strings.ToLower(input)
		if strings.HasPrefix(inputStr, "c") {
			return true
		}
		if strings.HasPrefix(inputStr, "h") {
			return false
		}
	}
}
