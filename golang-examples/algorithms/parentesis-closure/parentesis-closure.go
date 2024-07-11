// If a parentesis opens it must close.
package main

import "fmt"

func main() {
	tests := []struct {
		parens string
		want   bool
	}{
		{``, true},
		{`()`, true},
		{`((()))`, true},
		{`(6 (3 (1)(2) 3)(4)(5) 6) (7)`, true},
		{`)`, false},
		{`(`, false},
		{`)()(`, false},
		{`(()`, false},
		{`((()`, false},
		{`())`, false},
		{`()))`, false},
		{`((())))(`, false},
	}

	for _, test := range tests {
		got := areParensBalanced(test.parens)
		if got != test.want {
			fmt.Printf("[FAIL] areParensBalanced(%q) = %t; want %t\n", test.parens, got, test.want)
		}
	}
}

func areParensBalanced(parens string) bool {
	counter := 0
	for _, v := range parens {
		switch {
		case v == '(':
			counter++
		case v == ')':
			counter--
			if counter < 0 {
				return false
			}
		}
	}
	return counter == 0
}
