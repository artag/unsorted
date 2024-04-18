package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"sort"
	"strconv"
	"strings"
)

const (
	Num_digits  = 3
	Max_guesses = 10
	Fermi       = "Fermi"
	Pico        = "Pico"
	Bagels      = "Bagels"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := os.Stdout
	errOut := os.Stderr
	continuePlay := true

	for continuePlay {
		continuePlay = run(in, out, errOut)
	}

	fmt.Println("Thanks for playing!")
}

func run(in *bufio.Reader, out io.Writer, errOut io.Writer) bool {
	if err := clearConsole(out); err != nil {
		fmt.Fprintln(errOut, err)
		os.Exit(1)
	}

	introMsg := getIntroMessage()
	fmt.Print(introMsg)

	secret := getSecretNumber()
	guesses := 1
	for guesses <= Max_guesses {
		fmt.Printf("Guess #%d:\n", guesses)

		guess := getInputGuess(in, out)
		clue := getGlues(guess, secret)
		fmt.Println(clue)

		if guess == secret {
			break
		}
		if guesses >= Max_guesses {
			fmt.Println("You ran out of guesses.")
			fmt.Printf("The answer was %q.\n", secret)
		}

		guesses++
	}

	println("Do you want to play again? (yes or no)")
	playAgain := getInputContinue(in, out)
	return playAgain
}

func getIntroMessage() string {
	str := fmt.Sprintf(
		"Bagels, a deductive logic game.\n\n"+
			"I am thinking of a %d-digit number with no repeated digits.\n"+
			"Try to guess what it is. Here are some clues:\n"+
			"When I say:    That means:\n"+
			"    Pico       One digit is correct but in the wrong position.\n"+
			"    Fermi      One digit is correct and in the right position.\n"+
			"    Bagels     No digit is correct.\n\n"+
			"I have thought up a number.\n"+
			"You have %d guesses to get it.\n",
		Num_digits, Max_guesses)
	return str
}

func getSecretNumber() string {
	numbers := getShuffleNumbers()
	res := ""

	for i := 0; i < Num_digits; i++ {
		n := numbers[i]
		str := strconv.Itoa(n)
		res = res + str
	}

	return res
}

func getShuffleNumbers() [10]int {
	numbers := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	cnt := len(numbers)
	for i := range numbers {
		j := rand.Intn(cnt)
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}

	return numbers
}

func getGlues(guess, secret string) string {
	if guess == secret {
		return "You got it!"
	}

	clues := make([]string, 0)
	for i, r := range guess {
		ch := string(r)
		if ch == string(secret[i]) {
			clues = append(clues, Fermi)
		} else if strings.ContainsRune(secret, r) {
			clues = append(clues, Pico)
		}
	}

	if len(clues) < 1 {
		return Bagels
	}

	sort.Strings(clues)
	res := strings.Join(clues, " ")

	return res
}

func getInputGuess(in *bufio.Reader, out io.Writer) string {
	for {
		fmt.Fprint(out, "> ")
		input, err := in.ReadString('\n')
		input = strings.TrimRight(input, "\r\n")
		if err != nil {
			msg := fmt.Sprintf("%s. Enter integer %d-digit number.\n", err.Error(), Num_digits)
			fmt.Fprint(out, msg)
			continue
		}

		numStr := strings.TrimSuffix(input, "\n")
		if len(numStr) != Num_digits {
			msg := fmt.Sprintf("Enter integer %d-digit number.\n", Num_digits)
			fmt.Fprint(out, msg)
			continue
		}

		_, err = strconv.Atoi(numStr)
		if err != nil {
			msg := fmt.Sprintf("Wrong input: %q. Enter integer %d-digit number.\n", numStr, Num_digits)
			fmt.Fprint(out, msg)
			continue
		}

		return numStr
	}
}

func getInputContinue(in *bufio.Reader, out io.Writer) bool {
	for {
		fmt.Fprint(out, "> ")
		input, err := in.ReadString('\n')
		input = strings.TrimRight(input, "\r\n")
		if err != nil {
			msg := fmt.Sprintf("%s. Enter 'yes' or 'no'.\n", err.Error())
			fmt.Fprint(out, msg)
			continue
		}

		numStr := strings.TrimSuffix(input, "\n")
		if len(numStr) < 1 {
			fmt.Fprint(out, "Enter 'yes' or 'no'.\n")
			continue
		}

		low := strings.ToLower(input)
		playAgain := strings.HasPrefix(low, "y")
		return playAgain
	}
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
