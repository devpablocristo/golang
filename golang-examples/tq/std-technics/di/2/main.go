package main

import "fmt"

/////////////////////////////////////
// Component 1: Person
/////////////////////////////////////

type PersonHandler interface {
	SavePerson(person Person) error
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

/////////////////////////////////////

/////////////////////////////////////
// Component 2: InMemory
/////////////////////////////////////

var inMemDB []*Person

type PersonInMemory struct {
	handler PersonHandler
}

func NewPersonInMemory(handler PersonHandler) *PersonInMemory {
	return &PersonInMemory{handler: handler}
}

func (i *PersonInMemory) SavePerson(person *Person) error {
	inMemDB = append(inMemDB, person)
	return nil
}

func (i PersonInMemory) GetAllPersons() ([]*Person, error) {
	return inMemDB, nil
}

/////////////////////////////////////

/////////////////////////////////////
// Component 3: Repository
/////////////////////////////////////

type PersonRepositoryHandler interface {
	SavePerson(person *Person) error
}

type PersonRepository struct {
	personHandler PersonRepositoryHandler
}

func NewPersonRepository(handler PersonRepositoryHandler) *PersonRepository {
	return &PersonRepository{handler}
}

func (r *PersonRepository) SavePerson(person *Person) error {
	fmt.Println(*person)
	err := r.personHandler.SavePerson(person)
	if err != nil {
		return err
	}
	return nil
}

/////////////////////////////////////
// Components implementation
/////////////////////////////////////

func init() {

	p1 := Person{
		Name: "Bob",
		Age:  12,
	}

	p2 := Person{
		Name: "Mary",
		Age:  22,
	}

	inMemDB = append(inMemDB, &p1)
	inMemDB = append(inMemDB, &p2)

}

func main() {

	p3 := Person{
		Name: "John",
		Age:  30,
	}

	p4 := Person{
		Name: "Kurt",
		Age:  27,
	}

	var inMemHan PersonHandler
	inMem := NewPersonInMemory(inMemHan)

	inMem.SavePerson(&p3)
	inMem.SavePerson(&p4)

	persons, _ := inMem.GetAllPersons()

	showElements(persons)

	personRepo := NewPersonRepository(inMem)
	personRepo.SavePerson(&p4)
}

func showElements(s []*Person) {
	fmt.Println("Slice length:", len(s))
	for _, v := range s {
		fmt.Println(*v)
	}
}
