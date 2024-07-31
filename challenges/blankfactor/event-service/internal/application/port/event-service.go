package port

import (
	"context"

	domain "github.com/devpablocristo/blankfactor/event-service/internal/domain"
)

//go:generate mockgen -source=./event-service.go -destination=../../mocks/event-service_mock.go -package=mocks
type EventService interface {
	CreateEvent(context.Context, *domain.Event) (*domain.Event, error)
	GetAllEvents(context.Context) ([]*domain.Event, error)
	GetOverlappingEvents(context.Context) ([][]*domain.Event, error)
}
