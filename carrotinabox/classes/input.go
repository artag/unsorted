package classes

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Input struct {
	out io.Writer
	in  *bufio.Reader
}

func NewInput(out io.Writer, in *bufio.Reader) *Input {
	return &Input{
		out: out,
		in:  in,
	}
}

func (i *Input) EnterName(msg string) string {
	for {
		fmt.Fprintf(i.out, "%s", msg)

		input, err := i.in.ReadString('\n')
		input = strings.TrimRight(input, "\r\n")
		if err != nil {
			continue
		}
		if len(input) < 1 {
			continue
		}

		return input
	}
}

func (i *Input) PressEnter(msg string) {
	fmt.Fprintln(i.out, msg)
	_, _ = i.in.ReadString('\n')
	fmt.Fprintln(i.out)
}

func (i *Input) EnterYesNo(msg string) bool {
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

		answer := strings.ToLower(input)
		if strings.HasPrefix(answer, "y") {
			return true
		}
		if strings.HasPrefix(answer, "n") {
			return false
		}
	}
}

func (i *Input) EnterHumanOrComputer(msg string) bool {
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

		answer := strings.ToLower(input)
		if strings.HasPrefix(answer, "h") {
			return true
		}
		if strings.HasPrefix(answer, "c") {
			return false
		}
	}
}
