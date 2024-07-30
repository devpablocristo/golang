package main

import "fmt"

func main() {
	n := 8
	s := "UDDDUDUU"

	fmt.Println(numeroValles(n, s))
}

func numeroValles(n int, s string) int {
	nivelMar := 0
	valles := 0
	u := 0
	d := 0
	estoyValle := false

	for i := 0; i < n; i++ {
		if string(s[i]) == "U" {
			nivelMar++
			u++
		} else {
			nivelMar--
			d++
		}

		if nivelMar < 0 {
			if !estoyValle {
				valles++
			}
			estoyValle = true
		} else {
			estoyValle = false
		}

	}

	return valles

}
