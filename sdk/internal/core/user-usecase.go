package core

import (
	"context"
	"fmt"

	"github.com/devpablocristo/golang/sdk/internal/core/user"
)

type UserUseCasePort interface {
	GetUser(context.Context, string) (*user.User, error)
	DeleteUser(context.Context, string) error
	ListUsers(context.Context) (*user.InMemDB, error)
	UpdateUser(context.Context, *user.User, string) error
	CreateUser(context.Context, *user.User) error
	PublishMessage(context.Context, string) (string, error)
}

type userUseCase struct {
	user user.RepositoryPort
}

func NewUserUseCase(r user.RepositoryPort) UserUseCasePort {
	return &userUseCase{
		user: r,
	}
}

func (u *userUseCase) GetUser(ctx context.Context, ID string) (*user.User, error) {
	user, err := u.user.GetUser(ctx, ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUseCase) DeleteUser(ctx context.Context, ID string) error {
	return u.user.DeleteUser(ctx, ID)
}

func (u *userUseCase) ListUsers(ctx context.Context) (*user.InMemDB, error) {
	db, err := u.user.ListUsers(ctx)
	return db, err
}

func (u *userUseCase) UpdateUser(ctx context.Context, usr *user.User, ID string) error {
	return u.user.UpdateUser(ctx, usr, ID)
}

func (u *userUseCase) CreateUser(ctx context.Context, ucs *user.User) error {
	return u.user.SaveUser(ctx, ucs)
}

func (u *userUseCase) PublishMessage(ctx context.Context, msg string) (string, error) {
	fmt.Println("Message:", msg)
	return msg, nil
}
