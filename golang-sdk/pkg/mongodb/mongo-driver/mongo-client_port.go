package mongodbdriver

import "go.mongodb.org/mongo-driver/mongo"

type MongoDBClientPort interface {
	DB() *mongo.Database
	Close()
}
