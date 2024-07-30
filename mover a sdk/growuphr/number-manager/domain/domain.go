package domain

import (
	"time"
)

type User struct {
	UUID      string    `json:"uuid"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}

type Number struct {
	UUID      string    `json:"uuid"`
	Number    int       `json:"number"`
	CreatedAt time.Time `json:"created_at"`
}

type ReservedNumber struct {
	UUID      string    `json:"uuid"`
	Number    *Number   `json:"number" binding:"required"`
	User      *User     `json:"user" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
}
