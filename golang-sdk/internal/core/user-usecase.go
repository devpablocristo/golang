package core

import (
	"context"

	usr "github.com/devpablocristo/qh/events/internal/core/user"
)

type UserUseCase struct {
	repo usr.RepositoryPort
}

func NewUserUseCase(r usr.RepositoryPort) UserUseCasePort {
	return &UserUseCase{
		repo: r,
	}
}

func (uc *UserUseCase) GetUser(ctx context.Context, id string) (usr.User, error) {
	return uc.repo.GetUser(ctx, id)
}
