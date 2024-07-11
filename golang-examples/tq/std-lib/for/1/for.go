package main

import "fmt"

func main() {

	// for
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// for continued
	sum := 1
	for sum < 100 {
		sum += sum
	}
	fmt.Println("sum", sum)

	// for like while
	sumb := 1
	for sumb < 1000 {
		sumb += sumb
	}
	fmt.Println("sumb", sumb)

	// infinite for
	x := 0
	for {
		// do something in a loop and break manually
		x++
		if x >= 10 {
			fmt.Println("infite loop break")
			break
		}
	}
}
