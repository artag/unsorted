package internal

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

type Console struct {
	in  *bufio.Reader
	out io.Writer
}

func NewConsole(in io.Reader, out io.Writer) *Console {
	return &Console{
		in:  bufio.NewReader(in),
		out: out,
	}
}

func (c *Console) Clear() error {
	switch runtime.GOOS {
	case "linux":
		return clearConsoleLinux(c.out)
	case "windows":
		return clearConsoleWindows(c.out)
	}

	msg := fmt.Sprintf("Not supported OS %q\n", runtime.GOOS)
	return errors.New(msg)
}

func (c *Console) PressEnter() {
	in := bufio.NewReader(os.Stdin)
	_, _ = in.ReadString('\n')
	fmt.Println()
}

func (c *Console) InputNumber() int {
	for {
		fmt.Fprint(c.out, "Enter the sum: ")
		input, err := c.in.ReadString('\n')
		input = strings.TrimRight(input, "\r\n")
		if err != nil {
			continue
		}

		numStr := strings.TrimSuffix(input, "\n")
		num, err := strconv.Atoi(numStr)
		if err != nil {
			continue
		}

		return num
	}
}

func (c *Console) DisplayLine(format string, a ...any) {
	c.Display(format, a...)
	fmt.Fprintf(c.out, "\n")
}

func (c *Console) Display(format string, a ...any) {
	fmt.Fprintf(c.out, format, a...)
}

func clearConsoleLinux(out io.Writer) error {
	cmd := exec.Command("clear")
	cmd.Stdout = out
	return cmd.Run()
}

func clearConsoleWindows(out io.Writer) error {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = out
	return cmd.Run()
}
