package classes

import (
	"errors"
	"fmt"
	"io"
	"os/exec"
	"runtime"
)

type Console struct {
	out io.Writer
}

func NewConsole(out io.Writer) *Console {
	return &Console{out: out}
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
