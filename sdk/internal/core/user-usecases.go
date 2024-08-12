package core

import (
	"context"
	"fmt"

	"github.com/devpablocristo/golang/sdk/internal/core/user"
	"github.com/devpablocristo/golang/sdk/internal/core/user/ports"
)

// UserUseCases define los casos de uso para el manejo de usuarios
type UserUseCases interface {
	GetUser(context.Context, string) (*user.User, error)
	DeleteUser(context.Context, string) error
	ListUsers(context.Context) (*user.InMemDB, error)
	UpdateUser(context.Context, *user.User, string) error
	CreateUser(context.Context, *user.User) error
	PublishMessage(context.Context, string) (string, error)
}

type userUseCases struct {
	user ports.Repository
}

// NewUserUseCases crea una nueva instancia de UserUseCases
func NewUserUseCases(r ports.Repository) UserUseCases {
	return &userUseCases{
		user: r,
	}
}

func (u *userUseCases) GetUser(ctx context.Context, ID string) (*user.User, error) {
	user, err := u.GetUser(ctx, ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUseCases) DeleteUser(ctx context.Context, ID string) error {
	return u.DeleteUser(ctx, ID)
}

func (u *userUseCases) ListUsers(ctx context.Context) (*user.InMemDB, error) {
	db, err := u.ListUsers(ctx)
	return db, err
}

func (u *userUseCases) UpdateUser(ctx context.Context, usr *user.User, ID string) error {
	return u.UpdateUser(ctx, usr, ID)
}

func (u *userUseCases) CreateUser(ctx context.Context, ucs *user.User) error {
	return u.CreateUser(ctx, ucs)
}

func (u *userUseCases) PublishMessage(ctx context.Context, msg string) (string, error) {
	fmt.Println("Message:", msg)
	return msg, nil
}
