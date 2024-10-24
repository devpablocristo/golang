package entities

import (
	"time"
)

type UserType string

const (
	PersonType  UserType = "person"  // The user is a person
	CompanyType UserType = "company" // The user is a company
)

type User struct {
	UUID        string      // Unique identifier for the user
	PersonUUID  *string     // Reference to the associated person (nullable)
	CompanyUUID *string     // Reference to the associated company (nullable)
	UserType    UserType    // Defines if the user is a person or a company
	Credentials Credentials // Credentials information (username and password hash)
	Roles       []Role      // List of roles associated with the user
	CreatedAt   time.Time   // User creation date
	LoggedAt    time.Time   // Last time the user logged in (optional)
	UpdatedAt   time.Time   // Date when the user was last updated
	DeletedAt   *time.Time  // Date when the user was deleted (optional)
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
	Username     string // Username
	PasswordHash string // Password hash
}
