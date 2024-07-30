package user

import (
	"time"
)

type User struct {
	UUID       string    `json:"uuid" form:"uuid" gorm:"primary_key"`
	Username   string    `json:"username" form:"username" binding:"required"`
	Password   string    `json:"password" form:"password" binding:"required"`
	Email      string    `json:"email" form:"email" binding:"required"`
	CreatedAt  time.Time `bson:"created_at" json:"created_at"`
	UpdartedAt time.Time `gorm:"-" bson:"updated_at" json:"updated_at,omitempty"`
	DeletedAt  time.Time `gorm:"-" bson:"deleted_at" json:"deleted_at,omitempty"`
}

// type Users []User

// func (user *User) Validate() *errors.RestErr {
// 	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
// 	if user.Email == "" {
// 		return errors.BadRequestError("invalid email addreess")
// 	}
// 	return nil
// }
