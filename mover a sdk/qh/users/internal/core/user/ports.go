package usr

import "context"

type RepositoryPort interface {
	GetUser(context.Context, string) (*User, error)
	DeleteUser(context.Context, string) error
	ListUsers(context.Context) (*InMemDB, error)
	UpdateUser(context.Context, *User, string) error
	CreateUser(context.Context, *User) error
}
