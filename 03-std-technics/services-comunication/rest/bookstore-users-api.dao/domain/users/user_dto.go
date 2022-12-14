package users

import (
	"strings"
	"time"

	"github.com/devpablocristo/bookstore_users_api.dao/utils/errors"
)

// User datos del usuario
type User struct {
	Id        int64     `bson:"_id,omitempty" json:"id,omitempty"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at,omitempty"`
}

// Users lista de usuarios
type Users []*User

func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.BadRequestError("invalid email addreess")
	}
	return nil
}
