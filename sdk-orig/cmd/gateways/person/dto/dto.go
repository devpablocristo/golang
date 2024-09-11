package dto

import "github.com/devpablocristo/golang/sdk/internal/core/person/entities"

type PersonRequest struct {
	Firstname     string `json:"firstname" binding:"required"`
	Lastname      string `json:"lastname" binding:"required"`
	Age           int    `json:"age" binding:"required,gte=0"`
	Gender        string `json:"gender" binding:"required,oneof=male female other"`
	IDC           int64  `json:"idc" binding:"required"`
	Qualification int    `json:"qualification" binding:"gte=1,lte=10"`
}

func (dto *PersonRequest) ToDomain() *entities.Person {
	return &entities.Person{
		Firstname: dto.Firstname,
		Lastname:  dto.Lastname,
		Age:       dto.Age,
		Gender:    dto.Gender,
		IDC:       dto.IDC,
		//Qualification: dto.Qualification,
	}
}
