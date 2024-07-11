package mongodbrepo

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	config "crudmongodb/pkg/config"
)

func ConnectMongoDB(ctx context.Context) (*mongo.Client, error) {
	connectionURI := getConfig()
	clientOptions := options.Client().ApplyURI(connectionURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	log.Println("Conexión exitosa a MongoDB!")

	return client, nil
}

func getConfig() string {
	mongoUser := config.MongoUser
	mongoPassword := config.MongoPassword
	mongoHost := config.MongoHost
	mongoProtocol := config.MongoProtocol

	mongoURI := fmt.Sprintf("%s://%s:%s@%s/", mongoProtocol, mongoUser, mongoPassword, mongoHost)

	return mongoURI
}
