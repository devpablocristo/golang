package patients

import "time"

type Patient struct {
	ID           string
	PersonalInfo person.Person
	CreatedAt    time.Time
}
