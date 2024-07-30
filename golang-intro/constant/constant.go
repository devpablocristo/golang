package main

import "fmt"

// declare constant
const (
	MaxAge = 30
	MinAge = 15
)

func main() {

	// declare constant
	const months = 12

	// use constant
	var age int = 30
	fmt.Println("Age in months ", months*age)

	// use constant
	fmt.Println("Max age", MaxAge)
	fmt.Println("Min Age", MinAge)

}
