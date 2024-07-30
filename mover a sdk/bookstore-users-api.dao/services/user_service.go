package services

/*
	Proceso de de datos para interacción.
*/
import (
	users "github.com/devpablocristo/bookstore_users_api.dao/domain/users"
	"github.com/devpablocristo/bookstore_users_api.dao/utils/errors"
)

func CreateUser(u users.User) (*users.User, *errors.RestErr) {
	rErr := u.Validate()
	if rErr != nil {
		return nil, rErr
	}

	rErr = u.Save()
	if rErr != nil {
		return nil, rErr
	}

	return &u, nil
}

/*
func GetUsers() (*users.Users, *errors.RestErr) {
	urs, err := userRepository.GetUsers()
	if err != nil {
		return nil, err
	}

	return urs, nil
}
*/

func GetUser(uId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: uId}
	rErr := result.Get()
	if rErr != nil {
		return nil, rErr
	}
	return result, nil
}

/*
func UpdateUser(u users.User, uId string) (*users.User, *errors.RestErr) {
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
*/
