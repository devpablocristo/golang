package domain

type Doctor struct {
	ID           string
	PersonalInfo person.Person
	Speciality   string
}
