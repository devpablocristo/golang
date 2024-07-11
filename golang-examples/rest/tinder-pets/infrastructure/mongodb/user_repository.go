package mongodb

import (
	"context"

	"github.com/luka385/tinder-pets/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{
		collection: db.Collection("users"),
	}
}

func (r *UserRepository) Create(user *domain.User) error {
	_, err := r.collection.InsertOne(context.Background(), user)
	return err
}

func (r *UserRepository) Update(user *domain.User) error {
	filter := bson.M{"id": user.ID}
	update := bson.M{"$set": user}
	_, err := r.collection.UpdateOne(context.Background(), filter, update)
	return err
}

func (r *UserRepository) Delete(id string) error {
	filter := bson.M{"id": id}
	_, err := r.collection.DeleteOne(context.Background(), filter)
	return err
}

func (r *UserRepository) GetByID(id string) (*domain.User, error) {
	filter := bson.M{"id": id}
	user := &domain.User{}
	err := r.collection.FindOne(context.Background(), filter).Decode(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetAll() ([]*domain.User, error) {
	var users []*domain.User
	cursor, err := r.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	if err := cursor.All(context.Background(), &users); err != nil {
		return nil, err
	}
	return users, nil
}
