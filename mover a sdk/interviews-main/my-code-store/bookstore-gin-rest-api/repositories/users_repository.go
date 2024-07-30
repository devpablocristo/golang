package user

/*
   Interacción con la fuente de datos (CRUD) (ej: BBDD).
*/

import (
	"context"
	"time"

	db "github.com/devpablocristo/interviews/bookstore-gin-rest-api/databases"
	user "github.com/devpablocristo/interviews/bookstore-gin-rest-api/models/users"
	errors "github.com/devpablocristo/interviews/bookstore-gin-rest-api/utils/errors"
	"gopkg.in/mgo.v2/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection = db.GetCollection("users")
var ctx = context.Background()

//var ctx = context.TODO()

func CreateUser(u user.User) (*user.User, *errors.RestErr) {
	_, err := collection.InsertOne(ctx, u)
	if err != nil {
		restErr := errors.BadRequestError("ERROR! Inserting new document.")
		return nil, restErr
	}
	return &u, nil
}

func GetUsers() (*user.Users, *errors.RestErr) {
	filter := bson.M{}
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		restErr := errors.BadRequestError("ERROR! Reading collection.")
		return nil, restErr
	}

	var urs user.Users
	for cur.Next(ctx) {
		var u user.User
		err := cur.Decode(&u)
		if err != nil {
			restErr := errors.BadRequestError("ERROR! Reading documents.")
			return nil, restErr
		}
		urs = append(urs, &u)
	}
	return &urs, nil
}

func GetUser(uId string) (*user.User, *errors.RestErr) {
	oid, _ := primitive.ObjectIDFromHex(uId)
	filter := bson.M{"_id": oid}

	var u user.User
	err := collection.FindOne(ctx, filter).Decode(&u)
	if err != nil {
		restErr := errors.BadRequestError("ERROR! Reading document.")
		return nil, restErr
	}

	return &u, nil
}

func UpdateUser(u user.User, uId string) (*user.User, *errors.RestErr) {
	var err error
	oid, _ := primitive.ObjectIDFromHex(uId)
	filter := bson.M{"_id": oid}

	l, _ := time.LoadLocation("America/Buenos_Aires")
	t := time.Now()

	now := t.In(l)

	update := bson.M{
		"$set": bson.M{
			"username":   u.Username,
			"password":   u.Password,
			"email":      u.Email,
			"updated_at": now,
		},
	}

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		restErr := errors.BadRequestError("ERROR! Trying to update.")
		return nil, restErr
	}

	updU, rErr := GetUser(uId)
	if rErr != nil {
		restErr := errors.BadRequestError("ERROR! Finding new insertion.")
		return nil, restErr
	}

	return updU, nil
}

func DeleteUser(uId string) (*int64, *errors.RestErr) {
	var err error
	var oid primitive.ObjectID

	oid, err = primitive.ObjectIDFromHex(uId)
	if err != nil {
		restErr := errors.BadRequestError("ERROR! Using the uId.")
		return nil, restErr
	}

	filter := bson.M{"_id": oid}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		restErr := errors.BadRequestError("ERROR! Try to delete document.")
		return nil, restErr
	}

	return &result.DeletedCount, nil
}

func GetIdLastInseted() (bson.M, *errors.RestErr) {

	opts := options.FindOne().SetSort(bson.M{"$natural": -1})

	var lastDocument bson.M
	err := collection.FindOne(ctx, bson.M{}, opts).Decode(&lastDocument)
	if err != nil {
		restErr := errors.BadRequestError("ERROR! Trying to get last inseted document.")
		return nil, restErr
	}
	//sId := lastrecord["_id"].(primitive.ObjectID).Hex()
	//fmt.Println(err, sId)

	return lastDocument, nil
}
