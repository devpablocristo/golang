package event

import (
	"context"

	usrdom "github.com/devpablocristo/qh/internal/users/domain"
)

type RepoPort interface {
	CreateEvent(context.Context, *Event) (*Event, error)
	DeleteEvent(context.Context, string) (*Event, error)
	HardDeleteEvent(context.Context, string) (*Event, error)
	UpdateEvent(context.Context, *Event, string) (*Event, error)
	ReviveEvent(context.Context, string) (*Event, error)
	GetEvent(context.Context, string) (*Event, error)
	GetAllEvents(context.Context) ([]Event, error)
	AddUserToEvent(context.Context, string, *usrdom.User) (*Event, error)
}

type UseCasePort interface {
	CreateEvent(context.Context, *Event) (*Event, error)
	DeleteEvent(context.Context, string) (*Event, error)
	HardDeleteEvent(context.Context, string) (*Event, error)
	UpdateEvent(context.Context, *Event, string) (*Event, error)
	ReviveEvent(context.Context, string) (*Event, error)
	GetEvent(context.Context, string) (*Event, error)
	GetAllEvents(context.Context) ([]Event, error)
	AddUserToEvent(context.Context, string, *usrdom.User) (*Event, error)
}

type ConfigMongoPort interface {
	GetMongoURL() string
	GetMongoDBName() string
	GetMongoCollectionName() string
}

type ConfigGinPort interface {
	GetHandlerPort() string
}
