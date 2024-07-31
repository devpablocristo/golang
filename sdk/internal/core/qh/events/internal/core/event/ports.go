package event

import (
	"context"
)

type RepositoryPort interface {
	CreateEvent(context.Context, *Event) error
}
