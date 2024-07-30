// function multiple returns

package main

import "fmt"

func foo() (string, string) {
	return "two", "values"
}

func main() {
	fmt.Println(foo())

}
