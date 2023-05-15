package mongodbsrv

import (
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"crudmongodb/internal/domain"
	port "crudmongodb/internal/service/ports"
	config "crudmongodb/pkg/config"
)

type MongoService struct {
	client *mongo.Client
	ctx    context.Context
}

func NewMongoService(cli *mongo.Client, ctx context.Context) port.Repo {
	return &MongoService{
		client: cli,
		ctx:    ctx,
	}
}

func (ms *MongoService) Disconnect() {
	ms.client.Disconnect(ms.ctx)
}

func (ms *MongoService) Create(listing domain.Listing) error {
	collection := ms.client.Database(config.DbName).Collection(config.CollName)

	// Convertir la estructura Listing en BSON
	bsonData, err := bson.Marshal(listing)
	if err != nil {
		log.Println(err)
		return err
	}

	// Convertir los datos BSON en bson.M
	var document bson.M
	err = bson.Unmarshal(bsonData, &document)
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = collection.InsertOne(ms.ctx, document)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (ms *MongoService) Read(filter bson.M) ([]bson.M, error) {
	collection := ms.client.Database(config.DbName).Collection(config.CollName)
	findOptions := options.Find().SetProjection(bson.M{"name": 1, "_id": 0})

	cursor, err := collection.Find(ms.ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ms.ctx)

	var results []bson.M
	for cursor.Next(ms.ctx) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, errors.New("no documents found")
	}
	return results, nil
}

func (ms *MongoService) Update(filter bson.M, update bson.M) error {
	collection := ms.client.Database(config.DbName).Collection(config.CollName)
	res, err := collection.UpdateOne(ms.ctx, filter, update)
	if err != nil {
		return err
	}
	if res.ModifiedCount == 0 {
		return errors.New("no document was updated")
	}
	return nil
}

func (ms *MongoService) Delete(filter bson.M) error {
	collection := ms.client.Database(config.DbName).Collection(config.CollName)
	res, err := collection.DeleteOne(ms.ctx, filter)
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return errors.New("no document was deleted")
	}
	return nil
}
