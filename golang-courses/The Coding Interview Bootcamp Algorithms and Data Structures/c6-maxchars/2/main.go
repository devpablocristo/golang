package main

import "fmt"

func main() {
	s := "hhhello123333333333"
	fmt.Println(maxChar(s))

}

func maxChar(s string) string {
	m := make(map[string]int)
	for _, c := range s {
		m[string(c)] = m[string(c)] + 1
	}

	n := 0
	r := ""
	for key, value := range m {
		if value > n {
			n = value
			r = key
		}
	}

	return r

}
