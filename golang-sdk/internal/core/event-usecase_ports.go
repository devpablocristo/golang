package core

import (
	"context"

	eve "github.com/devpablocristo/qh/events/internal/core/event"
)

type EventUseCasePort interface {
	CreateEvent(context.Context, *eve.Event) error
}
