package main_test

import (
	"testing"
)

func TestRoundTime(t *testing.T) {

	t.Run("2h1m3s to 2h2m", func(t *testing.T) {
		week := 0
		day := 0
		hour := 2
		min := 1
		sec := 3

		wWeek := 0
		wDay := 0
		wHour := 2
		wMin := 2
		wSec := 0

		gWeek, gDay, gHour, gMin, gSec := ct.RoundTime(week, day, hour, min, sec)

		if wWeek != gWeek {
			t.Errorf("Week want: %d - Week got: %d", wWeek, gWeek)
		}
		if wDay != gDay {
			t.Errorf("Day want: %d - Day got: %d", wDay, gDay)

		}
		if wHour != gHour {
			t.Errorf("Hour want: %d - Hour got: %d", wHour, gHour)

		}
		if wMin != gMin {
			t.Errorf("Min want: %d - Min got: %d", wMin, gMin)

		}
		if wSec != gSec {
			t.Errorf("Sec want: %d - Sec got: %d", wSec, gSec)

		}

		t.Logf("Test Passed: %d,%d,%d,%d,%d", gWeek, gDay, gHour, gMin, gSec)

	})

}
