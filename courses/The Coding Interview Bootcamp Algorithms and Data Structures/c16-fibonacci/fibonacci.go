// Fibonacci
// The Fibonacci sequence is a set of integers (the Fibonacci numbers) that starts with a zero,
// followed by a one, then by another one, and then by a series of steadily increasing numbers.
// The sequence follows the rule that each number is equal to the sum of the preceding two numbers.
// In mathematics, the Fibonacci numbers, commonly denoted Fn, form a sequence, the Fibonacci sequence,
// in which each number is the sum of the two preceding ones.
// The sequence commonly starts from 0 and 1,
// although some authors omit the initial terms and start the sequence from 1 and 1 or from 1 and 2.
// Starting from 0 and 1, the Fibonacci sequence begins with the following 14 integers:
// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233 ...

package main

import "fmt"

func main() {

	fmt.Printf("\nfiboLoop: ")

	fmt.Println(fiboLoop(10))
	fmt.Println("----------------")

	fmt.Printf("fiboRecursion: ")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d, ", fiboRecursion(i))
	}
	fmt.Println("\n----------------")

	fmt.Printf("fiboClosure: ")
	f := fiboClosure()
	for i := 0; i < 10; i++ {
		fmt.Printf("%d, ", f())
	}
	fmt.Println("\n----------------")
}

// returns Array
func fiboLoop(n int) []int {
	// first 2 ints are added manually
	r := []int{0, 1}
	for i := 2; i < n; i++ {
		// here i starts with 2 so fn = fn-1 + fn-2 , r = (2-1) + (2-2) = 1
		r = append(r, r[i-1]+r[i-2])
	}
	return r
}

// closure function
func fiboClosure() func() int {
	// first 3 ints are added manually
	current, next, nextnext := 0, 1, 1

	return func() int {
		ret := current

		current = next
		next = nextnext
		nextnext = current + next
		return ret
	}
}

// Recursive function
func fiboRecursion(n int) int {
	if n < 2 {
		return n
	}
	return fiboRecursion(n-1) + fiboRecursion(n-2)
}
