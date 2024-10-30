package datamod

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UUID           uuid.UUID  `json:"uuid" db:"uuid"`
	PersonUUID     *uuid.UUID `json:"person_uuid,omitempty" db:"person_uuid"`
	Username       string     `json:"username" db:"username"`
	Password       string     `json:"password" db:"password"`
	EmailValidated bool       `json:"email_validated" db:"email_validated"`
	CreatedAt      time.Time  `json:"created_at" db:"created_at"`
	LoggedAt       *time.Time `json:"logged_at,omitempty" db:"logged_at"`
	UpdatedAt      *time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}
