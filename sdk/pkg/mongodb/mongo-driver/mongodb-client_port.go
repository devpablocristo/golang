package mongodbdriver

import "go.mongodb.org/mongo-driver/mongo"

// MongoDBClientPort é uma interface que define os métodos do cliente MongoDB
type MongoDBClientPort interface {
	DB() *mongo.Database
	Close()
}
