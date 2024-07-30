package person

type PersonDTO struct {
	Name          string
	Email         string
	Qualification int `validate:"gte=1,lte=10"`
}

func dtoToDomain(u *PersonDTO) Person {
	return Person{
		Name:          u.Name,
		Email:         u.Email,
		Qualification: u.Qualification,
	}
}
