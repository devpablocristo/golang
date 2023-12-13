package repo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type PortMongoGreetDAO interface{}
type MongoDBServicePort interface {
	Connect(ctx context.Context) (err error)
	Disconnect(ctx context.Context) error
	GetDatabase(ctx context.Context) (*mongo.Database, error)
	GetCollection(ctx context.Context) *mongo.Collection
}

type mongoGreetDAO struct {
	mongoService MongoDBServicePort
}

func NewMongoGreetDAO(ms MongoDBServicePort) PortMongoGreetDAO {
	return &mongoGreetDAO{
		mongoService: ms,
	}
}
