package main

import "fmt"

func main() {

	// map literal
	var x = map[int]string{
		1: "coffee",
		2: "te",
		3: "chocolate",
	}

	// adding more elements
	x[4] = "cappuccino"
	fmt.Println("key value of map x", x)

	// retrieve one element
	product := x[2]
	fmt.Println("Product of key 2", product)

	// delete element
	delete(x, 2)

	// check if element exist
	element, okExist := x[1]
	if okExist {
		fmt.Println("Element with id 1 is", element)
	}

	// create a map using make
	m := make(map[string]string)
	m["key1"] = "value1"
	m["key2"] = "value2"
	fmt.Println("key value of map m", m)

}
