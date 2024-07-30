package main

import "fmt"

func main() {
	s := "hello world"

	fmt.Println(reverse1(s))
	fmt.Println(reverse2(s))

}

func reverse1(s string) string {
	bs := []byte(s)

	for i, j := 0, len(bs)-1; i < j; i, j = i+1, j-1 {
		bs[i], bs[j] = bs[j], bs[i]
	}

	return string(bs)
}

func reverse2(s string) string {
	var rs string

	for i := len(s) - 1; i >= 0; i-- {
		rs += string(s[i])
	}

	return rs
}
