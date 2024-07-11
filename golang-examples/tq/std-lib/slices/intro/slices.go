package main

import "fmt"

func main() {

	// slide without data
	var colors []string
	fmt.Println("Colors without data", colors)

	colors = append(colors, "red")
	fmt.Println("Colors data", colors)

	// slice
	cities := []string{"buenos aires", "sao paulo", "new york"}
	fmt.Println("Cities", cities)

	// slice change value
	cities[1] = "bogota"
	fmt.Println("Cities with new value", cities)

	// slice using make, len and cap = 5
	m := make([]int, 5)
	fmt.Printf("Slice m - lenght=%d capacity=%d values %v", len(m), cap(m), m)
	fmt.Println()

	// slice using make, len = 0, cap = 5
	mc := make([]int, 0, 10)
	fmt.Printf("Slice mc - lenght=%d capacity=%d values %v", len(mc), cap(mc), mc)
	fmt.Println()

	mc = append(mc, 71)
	fmt.Println(mc)
}
