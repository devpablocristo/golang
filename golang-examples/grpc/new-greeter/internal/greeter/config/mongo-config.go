package config

import (
	"os"

	ctypes "github.com/devpablocristo/qh/internal/platform/custom-types"
)

type ConfigMongo struct {
	mongoURL            string
	mongoDBName         string
	mongoCollectionName string
}

func NewMongoConfig() ctypes.ConfigMongoPort {
	return &ConfigMongo{
		mongoURL:            os.Getenv("MONGO_URL"),
		mongoDBName:         os.Getenv("MONGO_GREETER_DB_NAME"),
		mongoCollectionName: os.Getenv("MONGO_GREETER_COLLECTION_NAME"),
	}
}

func (c *ConfigMongo) GetMongoURL() string {
	return c.mongoURL
}

func (c *ConfigMongo) GetMongoDBName() string {
	return c.mongoDBName
}

func (c *ConfigMongo) GetMongoCollectionName() string {
	return c.mongoCollectionName
}
