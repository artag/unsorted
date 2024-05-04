package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	i "calendarmaker/internal"
)

const ()

var (
	in = bufio.NewReader(os.Stdin)
)

func main() {
	year := enterYear()
	month := enterMonth()
	days := i.NewDays(month, year)
	calendar := i.NewCalendar(month, year, days)
	displayCalendar(calendar)
	if err := saveCalendarToFile(month, year, calendar); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func enterYear() int {
	for {
		fmt.Println("Enter the year for the calendar:")
		fmt.Print("> ")

		input, err := in.ReadString('\n')
		input = strings.TrimRight(input, "\r\n")
		if err != nil {
			continue
		}

		inputStr := strings.TrimSuffix(input, "\n")
		num, err := strconv.Atoi(inputStr)
		if err != nil {
			continue
		}
		if num < 0 {
			continue
		}

		return num
	}
}

func enterMonth() int {
	for {
		fmt.Println("Enter the month for the calendar, 1-12:")
		fmt.Print("> ")

		input, err := in.ReadString('\n')
		input = strings.TrimRight(input, "\r\n")
		if err != nil {
			continue
		}

		inputStr := strings.TrimSuffix(input, "\n")
		num, err := strconv.Atoi(inputStr)
		if err != nil {
			continue
		}
		if num < 1 || num > 12 {
			continue
		}

		return num
	}
}

func displayCalendar(calendar *i.Calendar) {
	fmt.Print(calendar.String())
}

func saveCalendarToFile(month, year int, calendar *i.Calendar) error {
	filename := fmt.Sprintf("calendar_%d_%d.txt", year, month)
	data := []byte(calendar.String())
	err := os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	fmt.Printf("Saved to %s\n", filename)
	return nil
}
