package main

import "fmt"

func Hello() string {
	return "Hello world!"
}

func Multiply(a, b int) int {
	r := a * b
	return r
}

func Hello2(s string) string {
	if s == "" {
		return "Hello world!"
	}

	return "Hello " + s + "!"

}

func main() {
	fmt.Println(Hello())
	fmt.Println(Hello2("Tori"))
}
