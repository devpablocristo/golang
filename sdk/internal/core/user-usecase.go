package core

import (
	"context"

	usr "github.com/devpablocristo/qh/events/internal/core/user"
)

// type UseCasePort interface {
// 	GetUser(context.Context, string) (*usr.User, error)
// 	DeleteUser(context.Context, string) error
// 	ListUsers(context.Context) (*usr.InMemDB, error)
// 	UpdateUser(context.Context, *usr.User, string) error
// 	CreateUser(context.Context, *usr.User) error
// 	PublishMessage(context.Context, string) (string, error)
// }

type UserUseCasePort interface {
	GetUser(context.Context, string) (usr.User, error)
}

type userUseCase struct {
	repo usr.RepositoryPort
}

func NewUserUseCase(r usr.RepositoryPort) UserUseCasePort {
	return &userUseCase{
		repo: r,
	}
}

func (uc *userUseCase) GetUser(ctx context.Context, id string) (usr.User, error) {
	return uc.repo.GetUser(ctx, id)
}

// type UseCase struct {
// 	usr usr.RepositoryPort
// }

// func NewUseCase(r usr.RepositoryPort) UseCasePort {
// 	return &UseCase{
// 		usr: r,
// 	}
// }

// func (u *UseCase) GetUser(ctx context.Context, ID string) (*usr.User, error) {
// 	user, err := u.usr.GetUser(ctx, ID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return user, nil
// }

// func (u *UseCase) DeleteUser(ctx context.Context, ID string) error {
// 	return u.usr.DeleteUser(ctx, ID)
// }

// func (u *UseCase) ListUsers(ctx context.Context) (*usr.InMemDB, error) {
// 	db, err := u.usr.ListUsers(ctx)
// 	return db, err
// }

// func (u *UseCase) UpdateUser(ctx context.Context, ucs *usr.User, ID string) error {
// 	return u.usr.UpdateUser(ctx, ucs, ID)
// }

// func (u *UseCase) CreateUser(ctx context.Context, ucs *usr.User) error {
// 	return u.usr.CreateUser(ctx, ucs)
// }

// func (u *UseCase) PublishMessage(ctx context.Context, msg string) (string, error) {
// 	fmt.Println("Message:", msg)
// 	return msg, nil
// }
