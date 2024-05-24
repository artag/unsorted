package internal

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
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

func (i *Input) PressEnter() {
	fmt.Fprintln(i.out, "Press 'Enter' to start countdown.")
	_, _ = i.in.ReadString('\n')
	fmt.Fprintln(i.out)
}

func (i *Input) EnterHours() int {
	min := 0
	max := 23
	return i.enterInteger(min, max, "Enter hours to countdown. From %d to %d:", min, max)
}

func (i *Input) EnterMinutes() int {
	min := 0
	max := 59
	return i.enterInteger(min, max, "Enter minutes to countdown. From %d to %d:", min, max)
}

func (i *Input) EnterSeconds() int {
	min := 0
	max := 59
	return i.enterInteger(min, max, "Enter seconds to countdown. From %d to %d:", min, max)
}

func (i *Input) enterInteger(min, max int, format string, a ...any) int {
	msg := fmt.Sprintf(format, a...)
	fmt.Fprintln(i.out, msg)
	for {
		fmt.Fprint(i.out, "> ")

		input, err := i.in.ReadString('\n')
		input = strings.TrimRight(input, "\r\n")
		if err != nil {
			continue
		}
		if len(input) < 1 {
			continue
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Fprintf(i.out, "Enter integer number from %d to %d.\n", min, max)
			continue
		}

		if min > num || num > max {
			fmt.Fprintf(i.out, "Enter integer number from %d to %d.\n", min, max)
			continue
		}

		return num
	}

}
