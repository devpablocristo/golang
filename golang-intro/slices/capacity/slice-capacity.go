package main

import (
	"fmt"
)

func main() {
	x := make([]int, 10, 12)

	fmt.Println(x)
	fmt.Println(len(x)) // logitud del slice
	fmt.Println(cap(x)) // capacidad del slice

	x = append(x, 333)

	fmt.Println(x)
	fmt.Println(len(x)) // logitud del slice
	fmt.Println(cap(x)) // capacidad del slice

	x = append(x, 1231)

	fmt.Println(x)
	fmt.Println(len(x)) // logitud del slice
	fmt.Println(cap(x)) // capacidad del slice

	// la capacidad es hasta donde se puede se pueden agregar elementos sin usar append
	// si se excede la capacidad, es necesario usar append para agregar un nuevo elemento y la capacidad si duplica
	// aunque se se duplica la capacidad, sigue siendo necesario usar append para agregar un nuevo elemento

	x = append(x, 99999)

	fmt.Println(x)
	fmt.Println(len(x)) // logitud del slice
	fmt.Println(cap(x)) // capacidad del slice

	x[5] = 12345

	fmt.Println(x)
	fmt.Println(len(x)) // logitud del slice
	fmt.Println(cap(x)) // capacidad del slice

}
