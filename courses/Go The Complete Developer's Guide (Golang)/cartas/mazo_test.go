package main

import (
	"fmt"
	"os"
	"testing"
)

func TestNuevoMazo(t *testing.T) {
	m := nuevoMazo()

	cant := 1234
	n := 4
	p := "Espadas"
	n2 := 6
	p2 := "Bastos"

	// esto es solo un ejemplo
	// implementar algo util mas adelante
	//if len(m) != 48 {
	if len(m) != cant {
		t.Errorf("Se esperaban %v cartas, pero hay %v.", cant, len(m))
	}

	fmt.Println(m[0].Numero)

	// no se pq no entra aqui
	//if m[0].Numero != 1 {
	if m[0].Numero != n {
		t.Errorf("Se esperaba una carta de numero %v, pero se obtuvo una de numero %v.", n, m[0].Numero)
	}

	//if m[0].Palo != "Copas" {
	if m[0].Palo != "Espadas" {
		t.Errorf("Se esperaba una carta de palo %v, pero se obtuvo de palo %v.", p, m[0].Palo)
	}

	if m[len(m)-9].Numero != n2 {
		t.Errorf("Se esperaba una carta de numero %v, pero se obtuvo una de numero %v.", n2, m[len(m)-9].Numero)
	}

	if m[len(m)-9].Palo != p2 {
		t.Errorf("Se esperaba una carta de palo %v, pero se obtuvo de palo %v.", p2, m[len(m)-9].Palo)
	}
}

/*
El nombre esta largo para que sea facil de encontrar
*/
func TestGuardaEnMazoYNuevoMazoDesdeArchivoCsv(t *testing.T) {
	os.Remove("_mazoTesting.csv")
	m := nuevoMazo()
	m.guardarEnArchivoCsv("_mazoTesting.csv")
	mazoCargado := nuevoMazoDesdeArchivoCsv("_mazoTesting.csv")
	n := 34
	if len(mazoCargado) != n {
		t.Errorf("Se esperaba len de %v, pero se obtuvo una de %v.", len(mazoCargado), n)
	}
	os.Remove("_mazoTesting.csv")
}
