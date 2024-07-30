package main

import (
	"fmt"
)

func hello(h string) string {
	/* logic*/
	return h
}

func sum(a, b int) (int, int, string) {
	s := a + b
	n := 7
	s2 := "oi!"
	return s, n, s2
}

func div(dv, ds float64) (float64, error) {
	if ds == 0 {

		/* logic */
		ds = 1
		//return 0, errors.New("dso must be positive")
	}

	res := dv / ds

	return res, nil
}

func main() {

	str := hello("hello world")

	fmt.Println(str)

	n1, n2, s1 := sum(1, 2)

	fmt.Println(n1, n2, s1)

	r, err := div(1, 2)
	if err != nil {
		fmt.Println("wrong division")
	}

	fmt.Println(r)

}
