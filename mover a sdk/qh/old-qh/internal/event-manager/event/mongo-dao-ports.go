package event

import (
	"context"

	usr "github.com/devpablocristo/qh/internal/user-manager/user"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoEventDAOPort interface {
	FindByID(context.Context, string) (*Event, error)
	Create(context.Context, *Event) (*Event, error)
	Update(context.Context, *Event, string) (*Event, error)
	HardDelete(context.Context, string) (*Event, error)
	List(context.Context) ([]Event, error)
	SoftDelete(context.Context, string) (*Event, error)
	SoftUndelete(context.Context, string) (*Event, error)
	AddUserToEvent(context.Context, string, *usr.User) (*Event, error)
}

type MongoDBServicePort interface {
	Connect(ctx context.Context) (err error)
	Disconnect(ctx context.Context) error
	GetDatabase(ctx context.Context) (*mongo.Database, error)
	GetCollection(ctx context.Context) *mongo.Collection
}
