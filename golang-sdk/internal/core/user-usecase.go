package core

import (
	"context"

	usr "github.com/devpablocristo/qh/events/internal/core/user"
)

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
