package main

import "fmt"

func sayHello() func(string) string {
	h := "YYY"
	return func(b string) string {
		h = h + " " + b
		return h
	}
}

func seqOfTwo() func() int {
	i := 0
	return func() int {
		i = i + 2
		return i
	}
}

func main() {
	a := sayHello()
	b := sayHello()

	fmt.Println(a("Hello golang"))
	fmt.Println(a("how are you?"))

	fmt.Println(b("Hi!"))
	fmt.Println(b("what up?"))

	c := seqOfTwo()
	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())

	d := seqOfTwo()
	fmt.Println(d())
	fmt.Println(d())
	fmt.Println(d())
}
