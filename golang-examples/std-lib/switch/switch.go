package main

import (
	"fmt"
	"time"
)

func main() {

	// basic switch - 'break' it is not necessary
	value := 4
	switch value {
	case 1:
		fmt.Println("value is one")
	case 2:
		fmt.Println("value is two")
	case 3:
		fmt.Println("value is three")
	case 4, 5:
		fmt.Println("value is four or five")
	default:
		fmt.Println("I do not recognize the value")
	}

	// switch without condition, like if
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Go home")
	}
}
