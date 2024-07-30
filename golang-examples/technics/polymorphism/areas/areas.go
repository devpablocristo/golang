package main

import (
	"fmt"
)

// Definimos la interfaz 'forma' que contiene dos métodos: 'area' y 'perimetro'.
// Cualquier tipo que implemente estos métodos será de tipo 'forma'.
type forma interface {
	area() float32
	perimetro() float32
}

// Definimos el tipo 'rectangulo' con dos campos: largo y ancho.
type rectangulo struct {
	largo float32
	ancho float32
}

// Definimos el tipo 'cuadrado' con un campo: lado.
type cuadrado struct {
	lado float32
}

// Definimos el tipo 'triangulo' con dos campos: base y altura.
type triangulo struct {
	base   float32
	altura float32
}

// Métodos del tipo 'rectangulo' que implementan la interfaz 'forma'.
// Calcula el área de un rectángulo.
func (r rectangulo) area() float32 {
	return r.largo * r.ancho
}

// Calcula el perímetro de un rectángulo.
func (r rectangulo) perimetro() float32 {
	return 2 * (r.largo + r.ancho)
}

// Método adicional del tipo 'rectangulo' que no está en la interfaz 'forma'.
// Este método no afecta la implementación de la interfaz.
func (r rectangulo) rotar() int {
	return 30
}

// Métodos del tipo 'cuadrado' que implementan la interfaz 'forma'.
// Calcula el área de un cuadrado.
func (c cuadrado) area() float32 {
	return c.lado * c.lado
}

// Calcula el perímetro de un cuadrado.
func (c cuadrado) perimetro() float32 {
	return 4 * c.lado
}

// Métodos del tipo 'triangulo' que implementan la interfaz 'forma'.
// Calcula el área de un triángulo isósceles.
func (t triangulo) area() float32 {
	return (t.base * t.altura) / 2
}

// Calcula el perímetro de un triángulo de forma simplificada.
func (t triangulo) perimetro() float32 {
	return 3 * t.base // Simplificación para fines de demostración.
}

// Función que imprime el área y el perímetro de cualquier tipo que implemente la interfaz 'forma'.
// Aquí es donde se demuestra el polimorfismo: podemos pasar cualquier tipo que implemente 'forma'.
func imprimirArea(f forma) string {
	return fmt.Sprintf("Area: %f - Perimetro: %f", f.area(), f.perimetro())
}

func main() {
	// Creamos instancias de cada tipo que implementa 'forma'.
	rec := rectangulo{
		largo: 10.0,
		ancho: 5.0,
	}

	cua := cuadrado{
		lado: 4.0,
	}

	tri := triangulo{
		base:   3.0,
		altura: 4.0,
	}

	// Usamos la función 'imprimirArea' para imprimir el área y perímetro de cada tipo.
	fmt.Println("Rectángulo: ", imprimirArea(rec))
	fmt.Println("Cuadrado: ", imprimirArea(cua))
	fmt.Println("Triángulo:", imprimirArea(tri))
}
