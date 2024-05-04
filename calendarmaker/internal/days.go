package internal

import "time"

type Days struct {
	days []int
}

func NewDays(month, year int) *Days {
	days := getDaysForCalendar(month, year)
	return &Days{
		days: days,
	}
}

func (d *Days) Pop() (bool, int) {
	if !d.Some() {
		return false, -1
	}

	day := d.days[0]
	d.days = d.days[1:]
	return true, day
}

func (d *Days) Some() bool {
	return len(d.days) > 0
}

func getDaysForCalendar(month, year int) []int {
	firstDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	beginDate := findBeginDate(firstDate)
	endDate := findEndDate(firstDate)
	return createDaysForCalendar(beginDate, endDate)
}

func findBeginDate(firstDate time.Time) time.Time {
	beginDate := firstDate
	for {
		if beginDate.Weekday() == time.Monday {
			break
		}
		beginDate = beginDate.AddDate(0, 0, -1)
	}
	return beginDate
}

func findEndDate(firstDate time.Time) time.Time {
	nextMonth := time.Month(firstDate.Month()%12 + 1)
	endDate := firstDate
	for {
		if endDate.Month() == nextMonth {
			break
		}
		endDate = endDate.AddDate(0, 0, 1)
	}

	endDate = endDate.AddDate(0, 0, -1)
	for {
		if endDate.Weekday() == time.Sunday {
			break
		}
		endDate = endDate.AddDate(0, 0, 1)
	}
	return endDate
}

func createDaysForCalendar(beginDate, end time.Time) []int {
	days := make([]int, 0)
	currDate := beginDate
	for currDate.Unix() <= end.Unix() {
		days = append(days, currDate.Day())
		currDate = currDate.AddDate(0, 0, 1)
	}

	return days
}
