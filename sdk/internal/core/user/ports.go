package user

import "context"

type RepositoryPort interface {
	SaveUser(context.Context, *User) error
	GetUser(context.Context, string) (*User, error)
	GetUserByUsername(context.Context, string) (*User, error)
	// DeleteUser(context.Context, string) error
	// ListUsers(context.Context) (*InMemDB, error)
	// UpdateUser(context.Context, *User, string) error
}
