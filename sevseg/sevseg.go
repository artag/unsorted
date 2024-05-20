package sevseg

import (
	"fmt"
	"strings"
)

var (
	validSymbols = map[rune]bool{
		'-': true,
		'.': true,
		'0': true,
		'1': true,
		'2': true,
		'3': true,
		'4': true,
		'5': true,
		'6': true,
		'7': true,
		'8': true,
		'9': true,
	}
)

func GetSevSegStr(number string, minWidth ...int) (string, error) {
	if err := checkNumber(number); err != nil {
		return "", err
	}

	num := getZeroPaddedNumber(number, minWidth)
	str := getNumberAsSegment(num)
	return str, nil
}

func checkNumber(number string) error {
	for _, ch := range number {
		_, contains := validSymbols[ch]
		if !contains {
			return fmt.Errorf("the number %q contains not valid symbol '%c'", number, ch)
		}
	}

	return nil
}

func getZeroPaddedNumber(number string, mw []int) string {
	minWidth := getMinWidthParam(mw)
	length := getMaxInt(len(number), minWidth)

	diff := length - len(number)
	num := number
	for diff > 0 {
		num = "0" + num
		diff--
	}

	return num
}

func getMinWidthParam(minWidth []int) int {
	var mw int

	if len(minWidth) < 1 {
		mw = 0
	} else {
		param := minWidth[0]
		if param < 0 {
			mw = 0
		} else {
			mw = param
		}
	}

	return mw
}

func getMaxInt(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}

	return num2
}

func getNumberAsSegment(num string) string {
	var row0 strings.Builder
	var row1 strings.Builder
	var row2 strings.Builder
	for i := 0; i < len(num); i++ {
		n := num[i]
		switch n {
		case '.':
			row0.WriteString(" ")
			row1.WriteString(" ")
			row2.WriteString(".")
		case '-':
			row0.WriteString("    ")
			row1.WriteString(" __ ")
			row2.WriteString("    ")
		case '0':
			row0.WriteString(" __ ")
			row1.WriteString("|  |")
			row2.WriteString("|__|")
		case '1':
			row0.WriteString("    ")
			row1.WriteString("   |")
			row2.WriteString("   |")
		case '2':
			row0.WriteString(" __ ")
			row1.WriteString(" __|")
			row2.WriteString("|__ ")
		case '3':
			row0.WriteString(" __ ")
			row1.WriteString(" __|")
			row2.WriteString(" __|")
		case '4':
			row0.WriteString("    ")
			row1.WriteString("|__|")
			row2.WriteString("   |")
		case '5':
			row0.WriteString(" __ ")
			row1.WriteString("|__ ")
			row2.WriteString(" __|")
		case '6':
			row0.WriteString(" __ ")
			row1.WriteString("|__ ")
			row2.WriteString("|__|")
		case '7':
			row0.WriteString(" __ ")
			row1.WriteString("   |")
			row2.WriteString("   |")
		case '8':
			row0.WriteString(" __ ")
			row1.WriteString("|__|")
			row2.WriteString("|__|")
		case '9':
			row0.WriteString(" __ ")
			row1.WriteString("|__|")
			row2.WriteString(" __|")
		}

		if i < len(num)-1 {
			// Non last character
			row0.WriteString(" ")
			row1.WriteString(" ")
			row2.WriteString(" ")
		} else {
			// The last character
			row0.WriteString("\n")
			row1.WriteString("\n")
			row2.WriteString("\n")
		}
	}

	str := row0.String() + row1.String() + row2.String()
	return str
}
