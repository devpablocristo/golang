package entities

import "time"

type Company struct {
	UUID      string     // Unique identifier for the company
	Cuit      string     // Company CUIT
	Name      string     // Company name
	LegalName string     // Legal name of the company (Razón Social)
	Address   string     // Company physical address
	Phone     string     // Company phone number
	Email     string     // Company email
	Industry  string     // Industry or sector of the company (e.g., Technology, Retail)
	CreatedAt time.Time  // Company creation date
	UpdatedAt time.Time  // Date when the company's info was last updated
	DeletedAt *time.Time // Date when the company was deleted (optional)
}
