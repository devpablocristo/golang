package main

import (
	"fmt"
	"strconv"
)

func NumOfSetBits(n int) int {
	count := 0
	for n != 0 {
		count += n & 1
		fmt.Println(count)
		n >>= 1
	}
	return count
}
func main() {
	n := 126
	fmt.Printf("Binary representation of %d is: %s.\n", n,
		strconv.FormatInt(int64(n), 2))
	fmt.Printf("The total number of set bits in %d is %d.\n", n, NumOfSetBits(n))
}
