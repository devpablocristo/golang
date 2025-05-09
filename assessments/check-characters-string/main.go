// Write a Golang code for checking if the given characters are present in a string.

package main

import (
	"fmt"
	"strings"
)

// Main function
func main() {

	// Create and initialize
	string1 := "Welcome to Interviewbit"
	string2 := "Golang Interview Questions"

	//  Check for presence Using Contains() method of strings package
	res1 := strings.Contains(string1, "Interview")
	res2 := strings.Contains(string2, "Go")

	// Displaying the result
	fmt.Println("Is 'Interview' present in string1 : ", res1)
	fmt.Println("Is 'Go' present in string2: ", res2)

}
