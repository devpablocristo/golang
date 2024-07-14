package main

import "fmt"

func printA() {
	fmt.Println("A")
}

func printB() {
	fmt.Println("B")
}

func printC() {
	fmt.Println("C")
}

func printD() {
	fmt.Println("D")
}

func main() {
	defer printA()
	defer printB()
	defer printC()
	defer printD()
	fmt.Println("Printing")
}
