package main

import "fmt"

func main() {
	si := []int{100, 200, 300, 100, 200, 400, 1, 453, 199, 1, 1, 1, 1, 1, 0}
	s := "nnn a a a a a n a"

	//fmt.Println(elements)

	// Test our method.
	sinEnterosDuplicados := eliminarEnterosDuplicados(si)
	fmt.Println(sinEnterosDuplicados)

	sinCaracteresDuplicados := eliminarCaracteresDuplicatesdos(s)
	fmt.Println(sinCaracteresDuplicados)

	result3 := contarAparicionCaracteres(s)
	fmt.Println(result3)

}

/*
elimina los enteros de duplicados de slice de enteros
*/
func eliminarEnterosDuplicados(si []int) []int {
	// enc map para marcar los enteros duplicados
	enc := map[int]bool{}
	r := []int{}

	for v := range si {
		if !enc[si[v]] {
			// Record this element as an encountered element.
			enc[si[v]] = true
			// Append to result slice.
			r = append(r, si[v])
		}
	}
	// Return the new slice.
	return r
}

/*
elimina los caracteres duplicados de un string
*/
func eliminarCaracteresDuplicatesdos(s string) string {
	// s a slice de bytes
	ss := []byte(s)
	// map para guardar los caracteres encontrados
	enc := map[byte]bool{}
	r := []byte{}

	for v := range ss {
		if !enc[ss[v]] {
			// Marcar como encontrado
			enc[ss[v]] = true
			// Append a r
			r = append(r, ss[v])
		}
	}
	// Return the new slice.
	return string(r)
}

func contarAparicionCaracteres(s string) map[string]int {
	bs := s
	enc := map[byte]bool{}
	r := []byte{}

	cantPorCaracter := map[string]int{}

	for v := range bs {
		// si ya fue encontrado sumar el número de encuentros
		if !enc[bs[v]] {
			cantPorCaracter[string(bs[v])]++
		} else {
			enc[bs[v]] = true
			cantPorCaracter[string(bs[v])] = 1
			r = append(r, bs[v])
		}
	}
	// Return the new slice.

	return cantPorCaracter
}
