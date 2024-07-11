// anonymous functions
package main

import "fmt"

func main() {

	a := "hello"
	x := func() string {
		fmt.Println(a)
		a = "bye"
		return a
	}() // here goes the value of the input parameter

	fmt.Println(a)
	fmt.Println(x)

	// this
	str := "Alice"
	str2 := "Jane"

	func(name string) {
		fmt.Println("Your name is", name)
	}(str2) // str2 = jane, so the inpunt param "name" is jane too

	//is same as:
	f := func(name string) {
		fmt.Println("Your name is", name)
	}
	f(str)
}
