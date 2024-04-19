package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	DaysInOneYear        = 365
	MaxDaysToAdd         = 364
	MinNumberOfBirthdays = 1
	MaxNumberOfBirthdays = 100
	SimulationTimes      = 100_000
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := os.Stdout

	clearConsole(out)
	fmt.Printf("How many birthdays shall I generate? (Max %d)\n", MaxNumberOfBirthdays)
	daysNumber := inputNumberOfBirthdays(in, out)

	fmt.Printf("Here are %d birthdays:\n", daysNumber)
	birthdays := getBirthdays(daysNumber)
	printDays(out, birthdays)

	match := getMatchDays(birthdays)
	printMatchDays(out, match)

	fmt.Printf("Generating %d random birthdays %d times...\n", daysNumber, SimulationTimes)
	waitToPressEnter(in, out, "Press Enter to begin...")

	fmt.Printf("Let's run another %d simulations.\n", SimulationTimes)
	simMatch := simulateTimes(out, daysNumber, SimulationTimes)
	fmt.Printf("%d simulations run.\n", SimulationTimes)

	probability := calculateProbability(simMatch, SimulationTimes)
	fmt.Printf(
		"Out of %d simulation of %d people, there was a\n"+
			"matching birthday in that group %d times. This means\n"+
			"that %d people have a %.2f%% chance of\n"+
			"having a matching birthday in their group.\n"+
			"That's probably more than you would think!\n",
		SimulationTimes, daysNumber, simMatch, daysNumber, probability)
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

func inputNumberOfBirthdays(in *bufio.Reader, out io.Writer) int {
	for {
		fmt.Fprint(out, "> ")
		input, err := in.ReadString('\n')
		input = strings.TrimRight(input, "\r\n")
		if err != nil {
			msg := fmt.Sprintf(
				"%s. Enter integer number (from %d to %d).\n",
				err.Error(), MinNumberOfBirthdays, MaxNumberOfBirthdays)
			fmt.Fprint(out, msg)
			continue
		}

		numStr := strings.TrimSuffix(input, "\n")
		num, err := strconv.Atoi(numStr)
		if err != nil {
			msg := fmt.Sprintf(
				"Wrong input: %q. Enter integer number (from %d to %d).\n",
				numStr, MinNumberOfBirthdays, MaxNumberOfBirthdays)
			fmt.Fprint(out, msg)
			continue
		}

		if num < MinNumberOfBirthdays || num > MaxNumberOfBirthdays {
			msg := fmt.Sprintf(
				"Wrong number: %d. Enter integer number (from %d to %d).\n",
				num, MinNumberOfBirthdays, MaxNumberOfBirthdays)
			fmt.Fprint(out, msg)
			continue
		}

		return num
	}
}

func getBirthdays(number int) []time.Time {
	initialDate := "2001-01-01"
	startOfYear, _ := time.Parse("2006-01-02", initialDate)

	birthdays := make([]time.Time, 0)
	for i := 0; i < number; i++ {
		days := rand.Intn(DaysInOneYear)
		birthday := startOfYear.AddDate(0, 0, days)
		birthdays = append(birthdays, birthday)
	}

	return birthdays
}

func printDays(out io.Writer, days []time.Time) {
	length := len(days)
	lastIdx := length - 1

	for idx, b := range days {
		var day string

		if idx == lastIdx {
			day = fmt.Sprintf("%s\n", b.Format("Jan 02"))
		} else {
			day = fmt.Sprintf("%s, ", b.Format("Jan 02"))
		}

		if idx != 0 && idx%8 == 0 {
			fmt.Fprintln(out)
		}
		fmt.Fprint(out, day)
	}

	fmt.Fprintln(out)
}

func waitToPressEnter(in *bufio.Reader, out io.Writer, message string) {
	fmt.Fprint(out, message)
	_, _ = in.ReadString('\n')
}

func printMatchDays(out io.Writer, match []time.Time) {
	fmt.Fprint(out, "In this simulation, ")

	if len(match) < 1 {
		fmt.Fprintf(out, "there are no matching birthdays.\n\n")
	} else {
		fmt.Fprintf(out, "multiple people have a birthday on:\n")
		printDays(out, match)
	}
}

func getMatchDays(birthdays []time.Time) []time.Time {
	length := len(birthdays)
	res := make([]time.Time, 0)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if birthdays[i] == birthdays[j] {
				res = append(res, birthdays[i])
			}
		}
	}
	return res
}

func simulateTimes(out io.Writer, daysNumber, simulateTimes int) int {
	simMatch := 0
	i := 0
	for i < simulateTimes {
		if i%10_000 == 0 {
			fmt.Fprintf(out, "%d simulations run...\n", i)
		}
		birthdays := getBirthdays(daysNumber)
		match := getMatchDays(birthdays)

		if len(match) > 0 {
			simMatch++
		}

		i++
	}

	return simMatch
}

func calculateProbability(simMatch, simTimes int) float64 {
	num := float64(simMatch) * 100 / float64(simTimes)
	probability := math.Round(num*100) / 100
	return probability
}
