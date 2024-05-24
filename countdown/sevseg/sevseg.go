package sevseg

import (
	"errors"
	"fmt"
	"strings"
)

var (
	validSymbols = map[rune]bool{
		' ': true,
		':': true,
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

func GetSegmentsInRow(numbers []string, minWidth int) (string, error) {
	if err := checkNumbers(numbers); err != nil {
		return "", err
	}

	for _, number := range numbers {
		if err := checkNumber(number); err != nil {
			return "", err
		}
	}

	var sb strings.Builder
	for _, number := range numbers {
		num := getZeroPaddedNumber(number, minWidth)
		sb.WriteString(num)
	}

	str := getNumberAsSegment(sb.String())
	return str, nil
}

func GetSegmentsInRows(numbers []string, minWidth int) ([]string, error) {
	if err := checkNumbers(numbers); err != nil {
		return make([]string, 0), err
	}

	for _, number := range numbers {
		if err := checkNumber(number); err != nil {
			return make([]string, 0), err
		}
	}

	var sb strings.Builder
	for _, number := range numbers {
		num := getZeroPaddedNumber(number, minWidth)
		sb.WriteString(num)
	}

	rows := getNumberAsSegmentAsRows(sb.String())
	return rows, nil
}

func checkNumbers(numbers []string) error {
	if len(numbers) < 1 {
		return errors.New("input numbers must not be empty")
	}

	return nil
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

func getZeroPaddedNumber(number string, mw int) string {
	if number == " " || number == ":" || number == "-" || number == "." {
		return number
	}

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

func getMinWidthParam(minWidth int) int {
	var mw int

	if minWidth < 0 {
		mw = 0
	} else {
		mw = minWidth
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
	rows := getNumberAsSegmentAsRows(num)
	var sb strings.Builder
	for _, row := range rows {
		sb.WriteString(row)
	}

	return sb.String()
}

func getNumberAsSegmentAsRows(num string) []string {
	var row0 strings.Builder
	var row1 strings.Builder
	var row2 strings.Builder
	for i := 0; i < len(num); i++ {
		n := num[i]
		switch n {
		case ' ':
			row0.WriteString("   ")
			row1.WriteString("   ")
			row2.WriteString("   ")
		case ':':
			row0.WriteString("   ")
			row1.WriteString(" * ")
			row2.WriteString(" * ")
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

	return []string{row0.String(), row1.String(), row2.String()}
}
