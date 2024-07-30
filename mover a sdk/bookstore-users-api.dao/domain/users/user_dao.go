package users

// DAO Pattern: Único punto de interaccion con la base de datos.

import (
	"fmt"

	"github.com/devpablocristo/bookstore_users_api.dao/datasources/mysql"
	"github.com/devpablocristo/bookstore_users_api.dao/utils/dates"
	"github.com/devpablocristo/bookstore_users_api.dao/utils/errors"
)

var (
	userDB = make(map[int64]*User)
)

func (u *User) Save() *errors.RestErr {
	result := userDB[u.Id]
	if result != nil {
		if result.Email == u.Email {
			return errors.BadRequestError(fmt.Sprintf("email %s already registered", u.Email))
		}
		return errors.NotFoundError(fmt.Sprintf("user %d already exists", u.Id))
	}

	u.CreatedAt = dates.GetNow()
	u.UpdatedAt = dates.GetNow()

	userDB[u.Id] = u
	return nil
}

/*
func GetUsers() (*user.Users, *errors.RestErr) {

	return &urs, nil
}
*/

func (u *User) Get() *errors.RestErr {
	err := mysql.UsersDB.Ping()
	if err != nil {
		panic(err)
	}

	result := userDB[u.Id]
	if result != nil {
		return errors.NotFoundError(fmt.Sprintf("user %d not found", u.Id))
	}

	u.Id = result.Id
	u.Username = result.Username
	u.Password = result.Password
	u.Email = result.Email
	u.CreatedAt = result.CreatedAt
	u.UpdatedAt = result.UpdatedAt

	return nil
}

/**
func UpdateUser(u user.User, uId string) (*user.User, *errors.RestErr) {

}

func DeleteUser(uId string) (*int64, *errors.RestErr) {

}

func GetIdLastInseted() (bson.M, *errors.RestErr) {

	return lastDocument, nil
}
*/
