package main

/*
	input: seconds as int
	output: string with max 2 units

	example:
	t= 7263
	output= 2h2m (rounded from 2h1m3s)
*/

import (
	"fmt"
)

func main() {
	t := 7263

	week, day, hour, min, sec := GetTimeUnitsValues(t)

	r := GetTimeFormeted(week, day, hour, min, sec)
	fmt.Println(r)
}

func GetTimeUnitsValues(t int) (int, int, int, int, int) {
	var week, day, hour, min, sec int = 0, 0, 0, 0, 0

	for t > 0 {
		switch {
		case t > 604800:
			week++
			t -= 604800
		case t > 86400 && t < 604800:
			day++
			t -= 86400
		case t > 3600 && t < 86400:
			hour++
			t -= 3600
		case t > 60 && t < 3600:
			min++
			t -= 60
		default:
			sec++
			t -= 1
		}
	}

	return week, day, hour, min, sec
}

func GetTimeFormeted(week, day, hour, min, sec int) string {
	var formatedTime string

	if week > 0 {
		formatedTime = fmt.Sprintf("%dw", week)
	}
	if day > 0 {
		formatedTime += fmt.Sprintf("%dd", day)
	}
	if hour > 0 {
		formatedTime += fmt.Sprintf("%dh", hour)
	}
	if min > 0 {
		formatedTime += fmt.Sprintf("%dm", min)
	}
	if sec > 0 {
		formatedTime += fmt.Sprintf("%ds", sec)
	}

	return formatedTime
}

func RoundTime(week, day, hour, min, sec int) (int, int, int, int, int) {
	units := 0
	if sec > 0 {
		units++
	}
	if min > 0 {
		units++
	}
	if min > 0 {
		units++
	}
	if day > 0 {
		units++
	}
	if week > 0 {
		units++
	}

	switch {
	case units > 4:
		day++
		hour = 0
		min = 0
		sec = 0
		return week, day, hour, min, sec
	case units > 3:
		hour++
		min = 0
		sec = 0
		return week, day, hour, min, sec
	case units > 2:
		min++
		sec = 0
		return week, day, hour, min, sec
	default:
		return week, day, hour, min, sec
	}
}
