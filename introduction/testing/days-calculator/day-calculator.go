package main

import (
	"time"
)

type Date struct {
	Day   int
	Month int
	Year  int
}

func main() {
	date1 := Date{
		Day:   31,
		Month: 4,
		Year:  1986,
	}
	date2 := Date{
		Day:   30,
		Month: 2,
		Year:  2345,
	}

	DaysCalculator(date1, date2)
}

func DaysCalculator(start, end Date) int {
	startDate := time.Date(start.Year, time.Month(start.Month), start.Day, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(end.Year, time.Month(end.Month), end.Day, 0, 0, 0, 0, time.UTC)
	days := int(endDate.Sub(startDate).Hours() / 24)
	return days
}
