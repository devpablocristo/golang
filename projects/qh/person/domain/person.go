package domain

import "time"

type Person struct {
	UUID      string    `json:"uuid" form:"uuid" gorm:"primary_key"`
	Firstname string    `json:"firstname" form:"firstname" binding:"required"`
	Lastname  string    `json:"lastname" form:"lastname" binding:"required"`
	Age       int       `json:"age" form:"age" binding:"required"`
	Gender    string    `json:"gender" form:"gender" binding:"required"`
	CreatedAt time.Time `gorm:"-" bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"-" bson:"updated_at" json:"updated_at,omitempty"`
	DeletedAt time.Time `gorm:"-" bson:"deleted_at" json:"deleted_at,omitempty"`
}
