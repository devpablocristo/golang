package person

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBServicePort interface {
	Connect(ctx context.Context) (err error)
	Disconnect(ctx context.Context) error
	GetDatabase(ctx context.Context) (*mongo.Database, error)
	GetCollection(ctx context.Context) *mongo.Collection
}
