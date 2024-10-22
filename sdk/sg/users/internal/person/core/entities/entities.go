package entities

import "time"

type Person struct {
	UUID        string     // Unique identifier for the person
	Cuil        string     // Person's CUIL
	Dni         string     // Person's DNI (optional)
	FirstName   string     // Person's first name
	LastName    string     // Person's last name
	Nationality string     // Person's nationality (optional)
	Email       string     // Person's email
	Phone       string     // Person's phone number
	CreatedAt   time.Time  // Person creation date
	UpdatedAt   time.Time  // Date when the person's info was last updated
	DeletedAt   *time.Time // Date when the person was deleted (optional)
}
