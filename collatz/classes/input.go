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

func (i *Input) EnterPositiveNumberOrQuit(msg string) (int, bool) {
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

		num, err := strconv.Atoi(input)
		if err != nil || num < 1 {
			fmt.Fprintln(i.out, "You must enter a number greater than 0.")
			continue
		}

		return num, false
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
