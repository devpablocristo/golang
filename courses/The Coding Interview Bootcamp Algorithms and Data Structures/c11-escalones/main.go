package main

import "fmt"

func main() {

	n := 4
	bs := []byte{}

	for i := 0; i < n; i++ {
		bs = append(bs, "#"...)
		fmt.Println(string(bs))
	}
}
