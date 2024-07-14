package main

import (
	"fmt"
)

// Interface
type IShape interface {
	area() float64
}

// Create a Rectangle struct
type Rectangle struct {
	Lenght float64
	Width  float64
}

// Create a Square struct
type Square struct {
	SideLength float64
}

// Create a Triangle struct
type Triangle struct {
	Base   float64
	Height float64
}

// Implement IShape.area() for Rectangle
func (r Rectangle) area() float64 {
	return r.Lenght * r.Width
}

// Implement IShape.area() for Square
func (s Square) area() float64 {
	return s.SideLength * 2
}

// Implement IShape.area() for Triangle
func (t Triangle) area() float64 {
	return t.Base * t.Height / 2
}

// Function to show Area
func ShowArea(s IShape) string {
	return fmt.Sprintf("Area %f", s.area())
}

func main() {
	// creating a Rectangle
	fmt.Println("Rectangle", ShowArea(Rectangle{10, 50}))
	fmt.Println("Square", ShowArea(Square{40}))
	fmt.Println("Triangle", ShowArea(Triangle{20, 30}))
}
