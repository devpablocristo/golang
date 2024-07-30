package main

import "fmt"

func main() {

	// create array and range from this array
	var numbers = []int{2, 4, 6, 8}
	for i, v := range numbers {
		fmt.Println("index of array", i, "value or array", v)
	}

	// create slice and range from this slice
	nums := []int{2, 4, 6}
	sum := 0
	for i, v := range nums {
		sum += v
		fmt.Println("index of slice", i, "sum", sum)
	}

	// create a map and range from this map
	fruits := map[string]string{"o": "orange", "b": "banana"}
	for k, v := range fruits {
		fmt.Println("key of map", k, "value of map", v)
	}

	// create string and range from this string
	s := "golan hello word"
	for i, c := range s {
		fmt.Println("index of character", i, "character", string(c))
	}

}
