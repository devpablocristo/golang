// The Animal struct contains Eat(), Sleep(), and Run() functions. 
// These functions are embedded into the child struct Dog by simply listing the struct at the top of the 
// implementation of Dog.



type Animal struct {
    // …
}

func (a *Animal) Eat()   { … }
func (a *Animal) Sleep() { … }
func (a *Animal) Run() { … }

type Dog struct {
    Animal
    // …
}