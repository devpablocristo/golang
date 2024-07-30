package main

import (
	"fmt"

	calc "encapsulation/calculator"
)

// command: go run main.go add.go sub.go
func main() {

	sum := calc.Add(2, 3)
	fmt.Println(sum)

	sub := calc.sub(3, 10)
	fmt.Println(sub)
}
