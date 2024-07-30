package user

import "context"

type RepositoryPort interface {
	Save(User) error
	FindByUsername(string) (User, error)
	GetUser(context.Context, string) (User, error)
	GetUserByUsername(context.Context, string) (User, error)
}
