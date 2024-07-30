package port

import (
	"context"

	domain "github.com/devpablocristo/golang/06-projects/qh/user/domain"
)

type Service interface {
	GetUsers(context.Context) (map[string]*domain.User, error)
	GetUser(context.Context, string) (*domain.User, error)
	CreateUser(context.Context, *domain.User) (*domain.User, error)
	UpdateUser(context.Context, string) error
	DeleteUser(context.Context, string) error
}
