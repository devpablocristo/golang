package user

import (
	"context"
	"fmt"

	entities "github.com/devpablocristo/golang/sdk/internal/core/user/entities"
	ports "github.com/devpablocristo/golang/sdk/internal/core/user/ports"
)

type useCases struct {
	repository ports.Repository
}

func NewUseCases(r ports.Repository) ports.UseCases {
	return &useCases{
		repository: r,
	}
}

func (u *useCases) GetUser(ctx context.Context, ID string) (*entities.User, error) {
	user, err := u.repository.GetUser(ctx, ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *useCases) GetUserUUID(ctx context.Context, username, passwordHash string) (string, error) {
	return "0001", nil
}

func (u *useCases) DeleteUser(ctx context.Context, ID string) error {
	return u.repository.DeleteUser(ctx, ID)
}

func (u *useCases) ListUsers(ctx context.Context) (*entities.InMemDB, error) {
	db, err := u.repository.ListUsers(ctx)
	return db, err
}

func (u *useCases) UpdateUser(ctx context.Context, usr *entities.User, ID string) error {
	return u.repository.UpdateUser(ctx, usr, ID)
}

func (u *useCases) CreateUser(ctx context.Context, ucs *entities.User) error {
	return u.repository.SaveUser(ctx, ucs)
}

func (u *useCases) PublishMessage(ctx context.Context, msg string) (string, error) {
	fmt.Println("Message:", msg)
	return msg, nil
}
