package main

import "fmt"

// command: go run main.go add.go sub.go
func main() {

	sum := add(2, 3)
	fmt.Println(sum)

	sub := sub(3, 10)
	fmt.Println(sub)
}
