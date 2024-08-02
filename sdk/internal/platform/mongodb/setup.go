package mongodb

import (
	"github.com/spf13/viper"

	mongodbdriver "github.com/devpablocristo/golang-sdk/pkg/mongodb/mongo-driver"
)

func NewMongoDBSetup() (mongodbdriver.MongoDBClientPort, error) {
	config := mongodbdriver.MongoDBClientConfig{
		User:     viper.GetString("MONGO_USER"),
		Password: viper.GetString("MONGO_PASSWORD"),
		Host:     viper.GetString("MONGO_HOST"),
		Port:     viper.GetString("MONGO_PORT"),
		Database: viper.GetString("MONGO_DATABASE"),
	}

	if err := config.Validate(); err != nil {
		return nil, err
	}

	if err := mongodbdriver.InitializeMongoDBClient(config); err != nil {
		return nil, err
	}

	return mongodbdriver.GetMongoDBInstance()
}
