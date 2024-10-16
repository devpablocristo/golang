package user

import (
	"context"
	"fmt"

	entities "github.com/devpablocristo/golang/sdk/qh/users/internal/core/entities"
	ports "github.com/devpablocristo/golang/sdk/qh/users/internal/core/ports"
)

type useCases struct {
	repository ports.Repository
}

func NewUseCases(r ports.Repository) ports.UseCases {
	return &useCases{
		repository: r,
	}
}

func (u *useCases) GetUserByUUID(ctx context.Context, UUID string) (*entities.User, error) {
	user, err := u.repository.GetUserByUUID(ctx, UUID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *useCases) GetUserByCredentials(ctx context.Context, username, passwordHash string) (string, error) {
	user, err := u.repository.GetUserByCredentials(ctx, username, passwordHash)
	if err != nil {
		return "", err
	}
	_ = user
	//return user, nil

	return "0001", nil
}

func (u *useCases) DeleteUser(ctx context.Context, UUID string) error {
	return u.repository.DeleteUser(ctx, UUID)
}

func (u *useCases) ListUsers(ctx context.Context) (*entities.InMemDB, error) {
	db, err := u.repository.ListUsers(ctx)
	return db, err
}

func (u *useCases) UpdateUser(ctx context.Context, usr *entities.User, UUID string) error {
	return u.repository.UpdateUser(ctx, usr, UUID)
}

func (u *useCases) CreateUser(ctx context.Context, ucs *entities.User) error {
	return u.repository.SaveUser(ctx, ucs)
}

func (u *useCases) PublishMessage(ctx context.Context, msg string) (string, error) {
	fmt.Println("Message:", msg)
	return msg, nil
}
