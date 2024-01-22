package main

import "fmt"

func printValue(value int) {
	fmt.Println("value argument with defer is", value)
}

func main() {
	v := 100
	defer printValue(v)
	v = 9999
	fmt.Println("value at the end is", v)
}
