package mongodb

import (
	"context"
	"errors"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	ctyoes "github.com/devpablocristo/qh/internal/platform/custom-types"
)

type MongoDBService struct {
	client *mongo.Client
	once   sync.Once

	mongoURL            string
	mongoDBName         string
	mongoCollectionName string
}

func NewMongoDBService(c ctyoes.ConfigMongoPort) *MongoDBService {
	return &MongoDBService{
		mongoURL:            c.GetMongoURL(),
		mongoDBName:         c.GetMongoDBName(),
		mongoCollectionName: c.GetMongoCollectionName(),
	}
}

func (m *MongoDBService) Connect(ctx context.Context) error {
	var err error
	m.once.Do(func() {
		log.Println("creating MongoDB connection pool")

		clientOptions := options.Client()
		clientOptions.ApplyURI(m.mongoURL)

		client, createErr := mongo.Connect(ctx, clientOptions)
		if createErr != nil {
			log.Println("mongo client error: ", err.Error())
			err = createErr
			return
		}

		m.client = client
		log.Println("connected to MongoDB")
	})

	return err
}

func (m *MongoDBService) Disconnect(ctx context.Context) error {
	if m.client == nil {
		return errors.New("MongoDB client is nil")
	}

	err := m.client.Disconnect(ctx)
	if err != nil {
		return err
	}

	log.Println("disconnected from MongoDB")
	return nil
}

func (m *MongoDBService) GetDatabase(ctx context.Context) (*mongo.Database, error) {
	if m.client == nil {
		return nil, errors.New("MongoDB client is not initialized")
	}
	return m.client.Database(m.mongoDBName), nil
}

func (m *MongoDBService) GetCollection(ctx context.Context) *mongo.Collection {
	if m.client == nil {
		log.Println("MongoDB client is not initialized")
		return nil
	}
	return m.client.Database(m.mongoDBName).Collection(m.mongoCollectionName)
}
