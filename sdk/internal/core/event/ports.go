package event

import (
	"context"
)

type RepositoryPort interface {
	CreateEvent(context.Context, *Event) error
	// DeleteEvent(context.Context, string) (*Event, error)
	// HardDeleteEvent(context.Context, string) (*Event, error)
	// UpdateEvent(context.Context, *Event, string) (*Event, error)
	// ReviveEvent(context.Context, string) (*Event, error)
	// GetEvent(context.Context, string) (*Event, error)
	// GetAllEvents(context.Context) ([]Event, error)
	// AddUserToEvent(context.Context, string, *usr.User) (*Event, error)
}

type DAOPort interface {
	Create(context.Context, *Event) error
	// FindByID(context.Context, string) (*Event, error)
	// Update(context.Context, *Event, string) (*Event, error)
	// HardDelete(context.Context, string) (*Event, error)
	// List(context.Context) ([]Event, error)
	// SoftDelete(context.Context, string) (*Event, error)
	// SoftUndelete(context.Context, string) (*Event, error)
	// AddUserToEvent(context.Context, string, *usr.User) (*Event, error)
}
