package main

import (
	"fmt"
	"math"
)

func main() {

	// Boolean
	var b bool = true
	if b {
		fmt.Println("is true")
	}

	// String
	var a string = "Hello World"
	fmt.Println(a)

	// Integer number
	var i int = 100
	fmt.Println(i + 99)

	// Float number
	var f float64 = 1.343435345563459
	fmt.Println(f)

	var fmax = math.MaxFloat64
	fmt.Println(fmax)

	// Complex number
	var c complex64 = 1.04i
	fmt.Println(c)

}
