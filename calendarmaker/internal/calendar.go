package internal

import (
	"fmt"
	"strings"
)

var (
	WeekDays = []string{
		"Понедельник", "Вторник", "Среда", "Четверг", "Пятница", "Суббота", "Воскресенье",
	}
	Months = []string{
		"Январь", "Февраль", "Март", "Апрель",
		"Май", "Июнь", "Июль", "Август",
		"Сентябрь", "Октябрь", "Ноябрь", "Декабрь",
	}
	WeekDaysCount = len(WeekDays)
	CellWidth     = getCellWidth()
	TableWidth    = CellWidth * WeekDaysCount
)

type Calendar struct {
	value string
}

func NewCalendar(month, year int, days *Days) *Calendar {
	calendar := createCalendar(month, year, days)
	return &Calendar{
		value: calendar,
	}
}

func (c *Calendar) String() string {
	return c.value
}

func createCalendar(month, year int, days *Days) string {
	var out strings.Builder

	printCaption(month, year, &out)
	printWeekDayNames(&out)
	printTable(days, &out)

	return out.String()
}

func printCaption(month int, year int, out *strings.Builder) {
	caption := fmt.Sprintf("%s %d", Months[month-1], year)
	captionLength := getStringLength(caption)
	captionIndent := (TableWidth - captionLength) / 2

	for i := 0; i < captionIndent; i++ {
		out.WriteRune(' ')
	}

	out.WriteString(fmt.Sprintf("%s\n\n", caption))
}

func printWeekDayNames(out *strings.Builder) {
	for i := 0; i < WeekDaysCount; i++ {
		day := fmt.Sprintf(" %s", WeekDays[i])
		out.WriteString(day)
		cnt := CellWidth - getStringLength(day)
		for j := 0; j < cnt+1; j++ {
			out.WriteString(" ")
		}
	}
	out.WriteString("\n")
}

func printTable(days *Days, out *strings.Builder) {
	for days.Some() {
		printTableHorizontalLine(out)
		printTableRowWithDates(out, days)

		for i := 0; i < 4; i++ {
			printTableRow(out)
		}
	}

	printTableHorizontalLine(out)
}

func printTableHorizontalLine(out *strings.Builder) {
	out.WriteString("+")
	for i := 0; i < WeekDaysCount; i++ {
		for j := 0; j < CellWidth; j++ {
			out.WriteString("-")
		}
		out.WriteString("+")
	}
	out.WriteString("\n")
}

func printTableRowWithDates(out *strings.Builder, days *Days) {
	out.WriteString("|")
	var width = CellWidth - 2
	for i := 0; i < WeekDaysCount; i++ {
		_, day := days.Pop()
		out.WriteString(fmt.Sprintf("%2d", day))

		for j := 0; j < width; j++ {
			out.WriteString(" ")
		}
		out.WriteString("|")
	}

	out.WriteString("\n")
}

func printTableRow(out *strings.Builder) {
	out.WriteString("|")
	for i := 0; i < WeekDaysCount; i++ {
		for j := 0; j < CellWidth; j++ {
			out.WriteString(" ")
		}
		out.WriteString("|")
	}

	out.WriteString("\n")
}

func getCellWidth() int {
	width := 0
	for _, day := range WeekDays {
		length := getStringLength(day)
		if width < length {
			width = length
		}
	}

	const additionalLength = 2
	width += additionalLength
	return width
}

func getStringLength(str string) int {
	symbols := strings.Split(str, "")
	length := len(symbols)
	return length
}
