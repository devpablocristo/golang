// Write a Go code to compare two slices of a byte.
package main

import (
	"bytes"
	"fmt"
)

func main() {

	sl1 := []byte{'I', 'N', 'T', 'E', 'R', 'V', 'I', 'E', 'W'}
	sl2 := []byte{'B', 'I', 'T'}

	// Use Compare function to compare slices
	res := bytes.Compare(sl1, sl2)

	if res == 0 {
		fmt.Println("Equal Slices")
	} else {
		fmt.Println("Unequal Slices")
	}
}
