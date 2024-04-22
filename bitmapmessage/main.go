package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const (
	bitmap = `
....................................................................
   **************   *  *** **  *      ******************************
  ********************* ** ** *  * ****************************** *
 **      *****************       ******************************
          *************          **  * **** ** ************** *
           *********            *******   **************** * *
            ********           ***************************  *
   *        * **** ***         *************** ******  ** *
               ****  *         ***************   *** ***  *
                 ******         *************    **   **  *
                 ********        *************    *  ** ***
                   ********         ********          * *** ****
                   *********         ******  *        **** ** * **
                   *********         ****** * *           *** *   *
                     ******          ***** **             *****   *
                     *****            **** *            ********
                    *****             ****              *********
                    ****              **                 *******   *
                    ***                                       *    *
                    **     *                    *
....................................................................
`
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := os.Stdout
	outErr := os.Stderr

	if err := clearConsole(out); err != nil {
		fmt.Fprintln(outErr, err)
		os.Exit(1)
	}

	displayWelcomeMessage(out)
	msg := enterMessage(in, out)
	displayMessageAsBitmap(msg)
}

func displayWelcomeMessage(out io.Writer) {
	fmt.Fprintln(out, "Bitmap message")
	fmt.Fprintln(out, "Enter the message to display with the bitmap.")
}

func enterMessage(in *bufio.Reader, out io.Writer) string {
	for {
		fmt.Fprint(out, "> ")
		input, err := in.ReadString('\n')
		input = strings.TrimRight(input, "\r\n")
		if err != nil {
			msg := fmt.Sprintf(
				"%s. Enter at least one character.\n", err.Error())
			fmt.Fprint(out, msg)
			continue
		}

		str := strings.TrimSuffix(input, "\n")
		if len(str) < 1 {
			fmt.Fprintf(out, "Enter at least one character.\n")
			continue
		}

		return str
	}
}

func displayMessageAsBitmap(message string) {
	msgLen := len(message)
	whiteSpace := rune(' ')

	lines := strings.Split(bitmap, "\n")

	for _, line := range lines {
		for i, char := range line {
			if char == whiteSpace {
				fmt.Print(" ")
			} else {
				ch := string(message[i%msgLen])
				fmt.Print(ch)
			}
		}
		fmt.Println()
	}

	fmt.Println()
}

func clearConsole(out io.Writer) error {
	switch runtime.GOOS {
	case "linux":
		return clearConsoleLinux(out)
	case "windows":
		return clearConsoleWindows(out)
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
