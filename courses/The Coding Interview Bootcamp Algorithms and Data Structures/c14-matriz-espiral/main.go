package main

import "fmt"

func main() {

	n := 3
	e := espiral(n)
	fmt.Println(e)
}

func espiral(n int) [][]int {

	r := [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}

	iniFil := 0
	iniCol := 0
	finFil := n - 1
	finCol := n - 1

	cont := 1
	for iniCol <= finCol && iniFil <= finFil {

		// primera fila
		for i := iniCol; i <= finCol; i++ {
			r[iniFil][i] = cont
			cont++
		}
		iniFil++

		// ultima fila
		for i := iniFil; i <= finFil; i++ {
			r[i][finCol] = cont
			cont++
		}
		finCol--

		//ultima fila
		for i := finCol; i >= iniCol; i-- {
			r[finFil][i] = cont
			cont++
		}
		finFil--

		//primera columna
		for i := finFil; i >= iniFil; i-- {
			r[i][iniCol] = cont
			cont++
		}
		iniCol++
	}

	return r
}
