package main

import (
	"fmt"
	"strconv"
)

func main() {

	for i := 0; i < 31; i++ {
		fmt.Println(fizzBuzzSwitch(i))
	}

	for i := 1; i < 31; i++ {
		fizzBuzzIf(i)
	}

}

func fizzBuzzSwitch(n int) string {
	switch {
	case n%3 == 0 && n%5 == 0:
		return "FizzBuzz"
	case n%3 == 0:
		return "Fizz"
	case n%5 == 0:
		return "Buzz"
	default:
		return strconv.Itoa(n)
	}
}

func fizzBuzzIf(n int) {
	if n%3 == 0 {
		if n%5 == 0 {
			fmt.Println("FizzBuzz")
		} else {
			fmt.Println("Fizz")
		}
	} else if n%5 == 0 {
		fmt.Println("Buzz")

	} else {
		fmt.Println(n)
	}
}
