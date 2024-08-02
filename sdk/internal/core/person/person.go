package person

import "time"

//import "github.com/go-playground/validator/v10"

type Person struct {
	UUID          string
	Firstname     string
	Lastname      string
	Age           int
	Gender        string
	IDC           int64
	Qualification int `validate:"gte=1,lte=10"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time //`gorm:"-" bson:"deleted_at" json:"deleted_at,omitempty"`
}
