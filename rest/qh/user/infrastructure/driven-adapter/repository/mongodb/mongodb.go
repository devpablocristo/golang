package mongodb

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	user     = "pablo"
	password = "toribio"
	database = "bookstore"
	host     = "cluster0.da7bu.mongodb.net"
	// port     = 1234 // en mongo atlas no hay difinicion de puerto
)

func GetCollection(collection string) *mongo.Collection {
	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority", user, password, host, database)
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	logErrors(err)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	logErrors(err)

	return client.Database(database).Collection(collection)
}

func logErrors(err error) {
	if err != nil {
		log.Fatal("ERROR!!! ", err)
	}
}
