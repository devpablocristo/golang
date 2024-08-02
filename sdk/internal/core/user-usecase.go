package core

import (
	"context"

	repo "github.com/devpablocristo/golang/sdk/internal/core/user"
)

type UserUseCasePort interface {
	GetUser(context.Context, string) (*repo.User, error)
	// DeleteUser(context.Context, string) error
	// ListUsers(context.Context) (*repo.InMemDB, error)
	// UpdateUser(context.Context, *repo.User, string) error
	// CreateUser(context.Context, *repo.User) error
	// PublishMessage(context.Context, string) (string, error)
}

type userUseCase struct {
	repo repo.RepositoryPort
}

func NewUserUseCase(r repo.RepositoryPort) UserUseCasePort {
	return &userUseCase{
		repo: r,
	}
}

func (u *userUseCase) GetUser(ctx context.Context, ID string) (*repo.User, error) {
	user, err := u.repo.GetUser(ctx, ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// func (u *userUseCase) DeleteUser(ctx context.Context, ID string) error {
// 	return u.repo.DeleteUser(ctx, ID)
// }

// func (u *userUseCase) ListUsers(ctx context.Context) (*repo.InMemDB, error) {
// 	db, err := u.repo.ListUsers(ctx)
// 	return db, err
// }

// func (u *userUseCase) UpdateUser(ctx context.Context, ucs *repo.User, ID string) error {
// 	return u.repo.UpdateUser(ctx, ucs, ID)
// }

// func (u *userUseCase) CreateUser(ctx context.Context, ucs *repo.User) error {
// 	return u.repo.CreateUser(ctx, ucs)
// }

// func (u *userUseCase) PublishMessage(ctx context.Context, msg string) (string, error) {
// 	fmt.Println("Message:", msg)
// 	return msg, nil
// }
