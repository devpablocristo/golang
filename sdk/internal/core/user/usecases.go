package user

import (
	"context"
	"fmt"

	"github.com/devpablocristo/golang/sdk/internal/core/user/entities"
	"github.com/devpablocristo/golang/sdk/internal/core/user/portscore"
)

type userUseCases struct {
	repository portscore.Repository
}

// NewUserUseCases crea una nueva instancia de UserUseCases
func NewUserUseCases(r portscore.Repository) portscore.UserUseCases {
	return &userUseCases{
		repository: r,
	}
}

func (u *userUseCases) GetUser(ctx context.Context, ID string) (*entities.User, error) {
	user, err := u.repository.GetUser(ctx, ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUseCases) DeleteUser(ctx context.Context, ID string) error {
	return u.repository.DeleteUser(ctx, ID)
}

func (u *userUseCases) ListUsers(ctx context.Context) (*entities.InMemDB, error) {
	db, err := u.repository.ListUsers(ctx)
	return db, err
}

func (u *userUseCases) UpdateUser(ctx context.Context, usr *entities.User, ID string) error {
	return u.repository.UpdateUser(ctx, usr, ID)
}

func (u *userUseCases) CreateUser(ctx context.Context, ucs *entities.User) error {
	return u.repository.SaveUser(ctx, ucs)
}

func (u *userUseCases) PublishMessage(ctx context.Context, msg string) (string, error) {
	fmt.Println("Message:", msg)
	return msg, nil
}
