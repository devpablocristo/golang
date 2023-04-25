package main

import "fmt"

func main() {

	// var string
	var hello string = "hello word"
	fmt.Println(hello)

	// var int
	var age int = 30
	fmt.Println("Age ", age)

	// var float
	var fb float64 = 30.12345678901234567890
	fmt.Println("Float fb ", fb)

	// multiple assignations
	var hi, word string = "hi", "word"
	fmt.Println(hi, word)

	// default value of variable int
	var defaultInt int
	fmt.Println("Value of defaultInt", defaultInt)

	// default value of boolean
	var defaultBoolean bool
	fmt.Println("Value of defaultBoolean", defaultBoolean)

	// short variable declarations
	great := "great work"
	fmt.Println(great)

	// type inference
	var b = true
	fmt.Println("Value of b", b)

}
