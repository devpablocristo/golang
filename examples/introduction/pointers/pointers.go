package main

import "fmt"

// value param
func addOne(value int) {
	value = value + 1
}

// pointer param, reference param
func addOnePointer(value *int) {
	*value = *value + 1
}

func main() {

	i := 1
	fmt.Println("Value of i is:", i)

	addOne(i)
	fmt.Println("addOne:", i)

	// using memory address of i (pointer)
	addOnePointer(&i)
	fmt.Println("addOnePointer:", i)

}
