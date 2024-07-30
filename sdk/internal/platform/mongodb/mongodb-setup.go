package mongodb

import (
	"github.com/spf13/viper"

	mongodbdriver "github.com/devpablocristo/qh/events/pkg/mongodb/mongo-driver"
)

// NewMongoDBSetup configura e inicializa uma conexão com MongoDB
func NewMongoDBSetup() (*mongodbdriver.MongoDBClient, error) {
	config := mongodbdriver.MongoDBClientConfig{
		User:     viper.GetString("MONGO_USER"),
		Password: viper.GetString("MONGO_PASSWORD"),
		Host:     viper.GetString("MONGO_HOST"),
		Port:     viper.GetString("MONGO_PORT"),
		Database: viper.GetString("MONGO_DATABASE"),
	}
	return mongodbdriver.NewMongoDBClient(config)
}
