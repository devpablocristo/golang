package domain

type Person struct {
	UUID     string
	Name     string
	Lastname string
	DNI      int64
	Gender   string
}

package person

//import "github.com/go-playground/validator/v10"

type Person struct {
	Name          string
	Email         string
	Qualification int `validate:"gte=1,lte=10"`
}
