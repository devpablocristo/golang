package event

import (
	"os"
)

type ConfigMongo struct {
	MongoURL            string
	MongoDBName         string
	MongoCollectionName string
}

func NewMongoConfig() ConfigMongoPort {
	return &ConfigMongo{
		MongoURL:            os.Getenv("MONGO_URL"),
		MongoDBName:         os.Getenv("MONGO_EVENTS_DB_NAME"),
		MongoCollectionName: os.Getenv("MONGO_EVENTS_COLLECTION_NAME"),
	}
}

func (c *ConfigMongo) GetMongoURL() string {
	return c.MongoURL
}

func (c *ConfigMongo) GetMongoDBName() string {
	return c.MongoDBName
}

func (c *ConfigMongo) GetMongoCollectionName() string {
	return c.MongoCollectionName
}
