package service

/*
	Proceso de de datos para interacción.
*/
import (
	"time"

	user "github.com/devpablocristo/interviews/bookstore-gin-rest-api/models/users"
	userRepository "github.com/devpablocristo/interviews/bookstore-gin-rest-api/repositories"
	"github.com/devpablocristo/interviews/bookstore-gin-rest-api/utils/errors"
)

func CreateUser(u user.User) (*user.User, *errors.RestErr) {
	rErr := u.Validate()
	if rErr != nil {
		return nil, rErr
	}

	l, _ := time.LoadLocation("America/Buenos_Aires")
	t := time.Now()

	u.CreatedAt = t.In(l)
	u.UpdatedAt = t.In(l)

	newU, err := userRepository.CreateUser(u)
	if err != nil {
		return nil, err
	}

	return newU, nil
}

func GetUsers() (*user.Users, *errors.RestErr) {
	urs, err := userRepository.GetUsers()
	if err != nil {
		return nil, err
	}

	return urs, nil
}

func GetUser(uId string) (*user.User, *errors.RestErr) {
	u, err := userRepository.GetUser(uId)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func UpdateUser(u user.User, uId string) (*user.User, *errors.RestErr) {
	ur, err := userRepository.UpdateUser(u, uId)
	if err != nil {
		return nil, err
	}

	return ur, nil
}

func DeleteUser(userId string) (*int64, *errors.RestErr) {
	del, err := userRepository.DeleteUser(userId)
	if err != nil {
		return nil, err
	}

	return del, nil
}
