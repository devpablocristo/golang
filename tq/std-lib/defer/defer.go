package main

import "fmt"

func sayGoodbye() {
	fmt.Println("Goodbye")
}

func sayHello() {
	defer sayGoodbye()
	fmt.Println("Hello")
}

func main() {
	sayHello()
}
