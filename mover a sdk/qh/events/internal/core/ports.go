package core

import (
	"context"

	eve "github.com/devpablocristo/qh/events/internal/core/event"
)

type UseCasePort interface {
	CreateEvent(context.Context, *eve.Event) error
}
