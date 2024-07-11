package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "hola que tal"

	fmt.Println(strings.Title(s))

	bs := s

	for _, v := range bs {
		if string(v) != " " {
			fmt.Println(string(v))
		}
	}

}
