package main

import "fmt"

func main() {
	// array tienen tamaño fijo
	// slice tienen tamaño dinámico
	a := [5]int{5, 6, 2, 7, 9}
	// para que funcione con copy hay que expresarlo asi
	si1 := make([]int, 5)
	si2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 91, 2, 3, 4, 5, 6, 7, 8, 91, 2, 3, 4, 5, 6, 7, 8, 91, 2, 3, 4, 5, 6, 7, 8, 91, 2, 3, 4, 5, 6, 7, 8, 91, 2, 3, 4, 5, 6, 7, 8, 91, 2, 3, 4, 5, 6, 7, 8, 9}

	// pasar desde array a slice
	// tb se puede de slice a array
	copy(si1, a[:5])

	chunk(si1, 2)
	chunk(si2, 5)
}

func chunk(si []int, tam int) {
	var siP [][]int
	var aux []int
	r := len(si) % tam

	if r > 0 {
		fmt.Println(r)
		aux = si[len(si)-r:]
		si = si[:len(si)-r]
	}

	ini := 0
	fin := tam

	for i := 0; i < len(si); i += tam {
		siP = append(siP, si[ini:fin])
		ini += tam
		fin += tam
	}

	if aux != nil {
		siP = append(siP, aux[:])
	}

	fmt.Println(siP)
}
