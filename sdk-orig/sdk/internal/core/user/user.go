package user

import (
	"time"
)

type User struct {
	UUID      string
	Username  string
	Password  string
	CreatedAt time.Time
}

type InMemDB map[string]*User
