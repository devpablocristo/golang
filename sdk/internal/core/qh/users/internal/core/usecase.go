package core

import (
	"context"
	"fmt"

	usr "github.com/devpablocristo/qh-users/internal/core/user"
)

type UseCase struct {
	usr usr.RepositoryPort
}

func NewUseCase(r usr.RepositoryPort) UseCasePort {
	return &UseCase{
		usr: r,
	}
}

func (u *UseCase) GetUser(ctx context.Context, ID string) (*usr.User, error) {
	user, err := u.usr.GetUser(ctx, ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UseCase) DeleteUser(ctx context.Context, ID string) error {
	return u.usr.DeleteUser(ctx, ID)
}

func (u *UseCase) ListUsers(ctx context.Context) (*usr.InMemDB, error) {
	db, err := u.usr.ListUsers(ctx)
	return db, err
}

func (u *UseCase) UpdateUser(ctx context.Context, ucs *usr.User, ID string) error {
	return u.usr.UpdateUser(ctx, ucs, ID)
}

func (u *UseCase) CreateUser(ctx context.Context, ucs *usr.User) error {
	return u.usr.CreateUser(ctx, ucs)
}

func (u *UseCase) PublishMessage(ctx context.Context, msg string) (string, error) {
	fmt.Println("Message:", msg)
	return msg, nil
}
