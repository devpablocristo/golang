package user

import (
	"time"

	"github.com/devpablocristo/golang/sdk/internal/core/auth"
	"github.com/devpablocristo/golang/sdk/internal/core/person"
)

type User struct {
	UUID          string
	Username      string
	Email         string
	Password      string
	Auth          *auth.Auth
	Person        *person.Person
	Qualification int `validate:"gte=1,lte=10"`
	CreatedAt     time.Time
	LoggedAt      time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}

type InMemDB map[string]*User
