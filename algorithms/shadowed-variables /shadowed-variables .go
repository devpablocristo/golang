package main

import "fmt"

func main() {
	i := 10 //scope: main
	j := 4
	for i := 'a'; i < 'b'; i++ {
		// i shadowed inside this block
		fmt.Println(i, j) //97 4
	}
	fmt.Println(i, j) //10 4

	if i := "test"; len(i) == j {
		// i shadowed inside this block
		fmt.Println(i, j) // i= test , j= 4
	} else {
		// i shadowed inside this block
		fmt.Println(i, j) //test 40
	}
	fmt.Println(i, j) //10 4
}
