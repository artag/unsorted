package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type InputData struct {
	Dices  int
	Sides  int
	Mod    int
	IsPlus bool
	Quit   bool
}

const (
	MinNumberOfDices = 1
	MaxNumberOfDices = 100
	MinNumberOfSides = 6
	MaxNumberOfSides = 38
)

var (
	in = bufio.NewReader(os.Stdin)
)

func main() {
	fmt.Println(`Dice Roller

Enter what kind and how many dice to roll. The format is the number of
dice, followed by "d", followed by the number of sides the dice have.
You can also add a plus or minus adjustment.

Example:
    3d6 rolls three 6-sided dice
    1d10+2 rolls one 10-sided dice, and adds 2
    2d38-1 rolls two 38-sided dice, and subtracks 1
    QUIT quits the program`)

	for {
		data := input()
		if data.Quit {
			fmt.Println("Thanks for playing!")
			os.Exit(0)
		}

		result := roll(data)
		fmt.Println(result)
	}
}

func input() InputData {
	for {
		fmt.Print("> ")

		input, err := in.ReadString('\n')
		input = strings.TrimRight(input, "\r\n")
		if err != nil {
			continue
		}
		if len(input) < 1 {
			continue
		}

		lin := strings.ToLower(input)
		if lin == "quit" {
			return InputData{Quit: true}
		}

		split1 := strings.Split(lin, "d")
		if len(split1) < 2 {
			fmt.Printf("Missing the 'd' character.\n")
			continue
		}

		numberOfDices, err := strconv.Atoi(split1[0])
		if err != nil {
			fmt.Printf("Input %q must be an integer number between %d and %d.\n", split1[0], MinNumberOfDices, MaxNumberOfDices)
			continue
		}

		if numberOfDices < MinNumberOfDices || numberOfDices > MaxNumberOfDices {
			fmt.Printf("Number of dices must between %d and %d.\n", MinNumberOfDices, MaxNumberOfDices)
			continue
		}

		isPlus := true
		split2 := strings.Split(split1[1], "+")
		if len(split2) < 2 {
			split2 = strings.Split(split1[1], "-")
			isPlus = false
		}

		numberOfSides, err := strconv.Atoi(split2[0])
		if err != nil {
			fmt.Printf("Input %q must be an integer number between %d and %d.\n", split2[0], MinNumberOfSides, MaxNumberOfSides)
			continue
		}

		if numberOfSides < MinNumberOfSides || numberOfSides > MaxNumberOfSides {
			fmt.Printf("Number of sides must between %d and %d.\n", MinNumberOfSides, MaxNumberOfSides)
			continue
		}

		mod := 0
		if len(split2) >= 2 {
			mod, err = strconv.Atoi(split2[1])
			if err != nil || mod < 0 {
				fmt.Printf("Input %q must be an positive integer number.\n", split2[1])
				continue
			}
		}

		return InputData{
			Dices:  numberOfDices,
			Sides:  numberOfSides,
			IsPlus: isPlus,
			Mod:    mod,
			Quit:   false,
		}
	}
}

func roll(data InputData) string {

	sum := 0
	scores := make([]string, 0)
	for i := 0; i < data.Dices; i++ {
		score := rand.Intn(data.Sides) + 1
		scores = append(scores, fmt.Sprint(score))
		sum += score
	}

	modStr := ""
	if data.IsPlus {
		sum += data.Mod
		if data.Mod > 0 {
			modStr = "+" + fmt.Sprint(data.Mod)
		}
	} else {
		sum -= data.Mod
		if data.Mod > 0 {
			modStr = "-" + fmt.Sprint(data.Mod)
		}
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprint(sum))
	sb.WriteString(" (")
	sb.WriteString(strings.Join(scores, ", "))
	if modStr != "" {
		sb.WriteString(", ")
		sb.WriteString(modStr)
	}
	sb.WriteRune(')')

	return sb.String()
}
