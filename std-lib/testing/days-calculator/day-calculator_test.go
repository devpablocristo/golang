package main_test

import (
	"testing"

	calc "github.com/devpablocristo/go-concepts/std-lib/testing/days-calculator"
)

func TestDaysCalculator(t *testing.T) {

	startDate := calc.Date{
		Day:   1,
		Month: 1,
		Year:  2020,
	}

	endDate := calc.Date{
		Day:   10,
		Month: 1,
		Year:  2020,
	}

	//startDate := time.Date(2020, time.Month(5), 1, 0, 0, 0, 0, time.UTC)
	//endDate := time.Date(2020, time.Month(5), 30, 0, 0, 0, 0, time.UTC)

	got := calc.DaysCalculator(startDate, endDate)
	want := 9

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}

}
