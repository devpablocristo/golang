package event

import (
	"context"

	entities "github.com/devpablocristo/golang/sdk/services/events-api/internal/event/entities"
)

type UseCases interface {
	CreateEvent(context.Context, *entities.Event) error
	// DeleteEvent(context.Context, string) (**entities.Event, error)
	// HardDeleteEvent(context.Context, string) (**entities.Event, error)
	// UpdateEvent(context.Context, **entities.Event, string) (**entities.Event, error)
	// ReviveEvent(context.Context, string) (**entities.Event, error)
	// GetEvent(context.Context, string) (**entities.Event, error)
	// GetAllEvents(context.Context) ([]*entities.Event, error)
	// AddUserToEvent(context.Context, string, *usr.User) (**entities.Event, error)
}
type Repository interface {
	CreateEvent(context.Context, *entities.Event) error
	// DeleteEvent(context.Context, string) (*entities.Event, error)
	// HardDeleteEvent(context.Context, string) (*entities.Event, error)
	// UpdateEvent(context.Context, *entities.Event, string) (*entities.Event, error)
	// ReviveEvent(context.Context, string) (*entities.Event, error)
	// GetEvent(context.Context, string) (*entities.Event, error)
	// GetAllEvents(context.Context) ([]Event, error)
	// AddUserToEvent(context.Context, string, *usr.User) (*entities.Event, error)
}

type DAO interface {
	Create(context.Context, *entities.Event) error
	// FindByID(context.Context, string) (*entities.Event, error)
	// Update(context.Context, *entities.Event, string) (*entities.Event, error)
	// HardDelete(context.Context, string) (*entities.Event, error)
	// List(context.Context) ([]Event, error)
	// SoftDelete(context.Context, string) (*entities.Event, error)
	// SoftUndelete(context.Context, string) (*entities.Event, error)
	// AddUserToEvent(context.Context, string, *usr.User) (*entities.Event, error)
}
