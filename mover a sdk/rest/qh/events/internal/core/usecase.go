package core

import (
	"context"

	eve "github.com/devpablocristo/qh/events/internal/core/event"
)

type UseCase struct {
	eve eve.RepositoryPort
}

func NewUseCase(r eve.RepositoryPort) UseCasePort {
	return &UseCase{
		eve: r,
	}
}

func (u *UseCase) CreateEvent(ctx context.Context, event *eve.Event) error {
	u.eve.CreateEvent(ctx, event)
	return nil
}
