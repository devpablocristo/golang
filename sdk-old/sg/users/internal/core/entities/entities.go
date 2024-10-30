package entities

import (
	"time"

	"github.com/google/uuid"

	datmod "github.com/devpablocristo/golang/sdk/sg/users/internal/adapters/connectors/data-model"
)

type UserType string

const (
	PersonType UserType = "person" // The user is a person
)

type User struct {
	UUID           string      // Unique identifier for the user
	PersonUUID     *string     // Reference to the associated person (nullable)
	Credentials    Credentials // Credentials information (username and password hash)
	EmailValidated bool
	Roles          []Role     // List of roles associated with the user
	CreatedAt      time.Time  // User creation date
	LoggedAt       *time.Time // Last time the user logged in (optional)
	UpdatedAt      *time.Time // Date when the user was last updated
	DeletedAt      *time.Time // Date when the user was deleted (optional)
}

type Role struct {
	Name        string       // Role name (e.g., Admin, User)
	Permissions []Permission // List of permissions associated with the role
}

type Permission struct {
	Name        string // Permission name (e.g., Create, Edit)
	Description string // Permission description
}

type Credentials struct {
	Username string // Username
	Password string // Password hash
}

func ToDataModel(user *User) (*datmod.User, error) {
	var dUser datmod.User
	var err error

	// Parse UUID from string to uuid.UUID
	dUser.UUID, err = uuid.Parse(user.UUID)
	if err != nil {
		return &dUser, err
	}

	// Parse PersonUUID if it's not nil
	if user.PersonUUID != nil {
		parsedPersonUUID, err := uuid.Parse(*user.PersonUUID)
		if err != nil {
			return &dUser, err
		}
		dUser.PersonUUID = &parsedPersonUUID
	} else {
		dUser.PersonUUID = nil
	}

	// Map Credentials
	dUser.Username = user.Credentials.Username
	dUser.Password = user.Credentials.Password

	dUser.EmailValidated = user.EmailValidated

	dUser.CreatedAt = user.CreatedAt
	dUser.LoggedAt = user.LoggedAt
	dUser.UpdatedAt = user.UpdatedAt
	dUser.DeletedAt = user.DeletedAt

	return &dUser, nil
}
