package port

import (
	"context"

	domain "github.com/devpablocristo/99minutos/order-manager/internal/domain"
)

//go:generate mockgen -source=./user-repo.go -destination=../../mocks/user-repo_mock.go -package=mocks
type UserRepo interface {
	Create(context.Context, *domain.User) error
	Read(context.Context, string) (*domain.User, error)
	List(context.Context) map[string]*domain.User
	Delete(context.Context, string) error
	FindByEmail(context.Context, string) (*domain.User, error)
}
