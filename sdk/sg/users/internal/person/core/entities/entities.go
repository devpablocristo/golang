package entities

type Person struct {
	UUID        string // Unique identifier for the person
	Cuil        string // Person's CUIL
	Dni         string // Person's DNI (optional)
	FirstName   string // Person's first name
	LastName    string // Person's last name
	Nationality string // Person's nationality (optional)
	Email       string // Person's email
	Phone       string // Person's phone number
}
