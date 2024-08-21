package domain

import "github.com/devpablocristo/golang/sdk/internal/core/person"

type Doctor struct {
	ID           string
	PersonalInfo person.Person
	Speciality   string
}
