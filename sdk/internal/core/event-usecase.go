package core

import (
	"context"

	eve "github.com/devpablocristo/qh/events/internal/core/event"
)

type eventUsecase struct {
	eve eve.RepositoryPort
}

func NewUseCase(r eve.RepositoryPort) EventUseCasePort {
	return &eventUsecase{
		eve: r,
	}
}

func (u *eventUsecase) CreateEvent(ctx context.Context, event *eve.Event) error {
	u.eve.CreateEvent(ctx, event)
	return nil
}
