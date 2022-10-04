package service

import (
	"fmt"

	ports "github.com/JoakoMeLi/go-concepts/3rd-party-libs/gomock/ports"
)

type Person struct {
	First string
}

// Accesor is how to store or retrieve a person
// When retriving a person, if they do not exist, return the zero value

type PersonService struct {
	a ports.Accesor
}

func NewPersonService(a ports.Accessor) PersonService {
	return PersonService{
		a: a,
	}
}

func (ps PersonService) Get(n int) (Person, error) {
	p := ps.a.Retrieve(n)
	if p.First == "" {
		return Person{}, fmt.Errorf("No person with n of %d", n)
	}
	return p, nil
}

func Put(a ports.Accessor, n int, p Person) {
	a.Save(n, p)
}

func Get(a ports.Accessor, n int) Person {
	return a.Retrieve(n)
}
