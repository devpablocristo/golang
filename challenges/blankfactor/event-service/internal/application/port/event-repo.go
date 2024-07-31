package port

import (
	"context"

	domain "github.com/devpablocristo/blankfactor/event-service/internal/domain"
)

//go:generate mockgen -source=./event-repo.go -destination=../../mocks/event-repo_mock.go -package=mocks
type EventRepo interface {
	CreateEvent(context.Context, *domain.Event) (*domain.Event, error)
	GetEventByID(context.Context, int64) (*domain.Event, error)
	UpdateEvent(context.Context, *domain.Event, int64) error
	DeleteEvent(context.Context, int64) error
	GetAllEvents(context.Context) ([]*domain.Event, error)
}
