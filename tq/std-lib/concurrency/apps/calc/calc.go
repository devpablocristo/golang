package main

import (
	"fmt"
)

func calcSum(number1 int, number2 int, chsum chan int) {
	sum := number1 + number2
	chsum <- sum
}

func calcMultiplication(number1 int, number2 int, chprod chan int) {
	prod := number1 * number2
	chprod <- prod
}

func main() {
	sumCh := make(chan int)
	prodCh := make(chan int)

	go calcSum(20, 30, sumCh)
	go calcMultiplication(40, 50, prodCh)

	sum, prod := <-sumCh, <-prodCh

	fmt.Println("Result sum", sum)
	fmt.Println("Result prod", prod)
}
