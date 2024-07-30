package mongodbrepo

import (
	"context"
	"errors"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	domain "crudmongodb/internal/domain"
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

	document, err := listingToBSON(listing)
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

func (ms *MongoService) ReadAll() ([]domain.Listing, error) {
	//var listings []domain.Listing

	listings := make([]domain.Listing, 0, 10000)

	collection := ms.client.Database(config.DbName).Collection(config.CollName)

	// parece que el problema es al agregar info, la agrego de cq forma eso hace no ande
	// borrra lo guardado y volver a probar traer la info desde adress pelado
	filter := bson.M{
		"name": 1, // Incluye el campo "name"
		//"address": 1, // Incluye el campo "address"
		// "address.street": 1, // Incluye el campo "address"
		// "address.market": 1,
		//"summary":        1, // Incluye el campo "summary"
		"_id": 1, // Excluye el campo "_id"
		//"address.location": 1, // Incluye el campo "address"
	}

	findOptions := options.Find().SetProjection(filter)

	// Se ejecuta la consulta Find en la colección para recuperar todos los documentos
	// sin aplicar ningún filtro. Se utiliza un filtro vacío bson.M{} y las opciones de
	// búsqueda findOptions.
	cursor, err := collection.Find(ms.ctx, bson.M{}, findOptions)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer cursor.Close(ms.ctx)

	// Se itera a través del cursor para obtener cada documento devuelto por la consulta.
	i := 0
	for cursor.Next(ms.ctx) {
		i++

		//for i := 0; i < 10000; i++ {
		cursor.Next(ms.ctx)
		var listing domain.Listing
		// Se decodifica el documento en un objeto Listing utilizando la función Decode del cursor.
		// El operador & se utiliza para pasar un puntero a la variable listing para que la decodificación
		// se realice correctamente.

		err := cursor.Decode(&listing)
		if err != nil {
			log.Println("Este es el error:", err)
			return nil, err
		}

		fmt.Println(listing)
		fmt.Println(i)
		// Se agrega el objeto Listing al slice listings utilizando append para almacenar
		// todos los documentos recuperados.
		listings = append(listings, listing)
	}

	// Se verifica si hubo algún error durante la iteración del cursor utilizando el método Err() del cursor.
	// Si hay un error, se retorna el error.
	if err := cursor.Err(); err != nil {
		log.Println(err)
		return nil, err
	}

	// Se verifica si no se encontraron documentos.
	// Si la longitud del slice listings es cero, se retorna un error indicando que no se encontraron documentos.
	if len(listings) == 0 {
		log.Println(err)
		return nil, errors.New("no documents found")
	}

	fmt.Println("listings")

	return listings, nil
}

func (ms *MongoService) ReadByID(filter bson.M) ([]bson.M, error) {
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

func bsonToListing(document bson.M) (domain.Listing, error) {
	bsonData, err := bson.Marshal(document)
	if err != nil {
		log.Println(err)
		return domain.Listing{}, err
	}

	var listing domain.Listing
	err = bson.Unmarshal(bsonData, &listing)
	if err != nil {
		log.Println(err)
		return domain.Listing{}, err
	}

	return listing, nil
}

func listingToBSON(listing domain.Listing) (bson.M, error) {
	bsonData, err := bson.Marshal(listing)
	if err != nil {
		log.Println(err)
		return bson.M{}, err
	}

	var document bson.M
	err = bson.Unmarshal(bsonData, &document)
	if err != nil {
		log.Println(err)
		return bson.M{}, err
	}

	return document, nil
}
