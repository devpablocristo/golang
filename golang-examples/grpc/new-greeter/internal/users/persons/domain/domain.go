package domain

//import "github.com/go-playground/validator/v10"

type Person struct {
	Name          string
	Email         string
	Qualification int `validate:"gte=1,lte=10"`
}
