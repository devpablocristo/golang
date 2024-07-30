package db

import (
	"context"
	"log"

	"github.com/vidu171/clean-architecture-go/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage struct {
	MongoClient mongo.Client
	database    *mongo.Database
}

func NewStorage(connectString string, dbName string) (Storage, error) {
	Storage := Storage{}
	clientOptions := options.Client().ApplyURI(connectString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return Storage, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return Storage, err
	}
	Storage.MongoClient = *client
	Storage.database = client.Database(dbName)
	return Storage, nil
}

func (Storage Storage) FindAllBooks() ([]*domain.Book, error) {
	var results []*domain.Book
	collection := Storage.database.Collection("books")
	cur, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cur.Next(context.TODO()) {
		var elem domain.Book
		err2 := cur.Decode(&elem)
		if err2 != nil {
			log.Fatal(err2)
		}
		results = append(results, &elem)
	}
	return results, nil
}

func (Storage Storage) SaveBook(book domain.Book) error {
	collection := Storage.database.Collection("books")
	_, err := collection.InsertOne(context.TODO(), book)
	if err != nil {
		return err
	}
	return nil
}
func (Storage Storage) SaveAuthor(author domain.Author) error {
	collection := Storage.database.Collection("authors")
	_, err := collection.InsertOne(context.TODO(), author)
	if err != nil {
		return err
	}
	return nil
}
