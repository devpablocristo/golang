package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"chat/internal/domain"
)

type MongoDB struct {
	db *mongo.Database
}

func NewRepository(db *mongo.Database) *MongoDB {
	return &MongoDB{
		db: db,
	}
}

func (r *MongoDB) SaveMessage(ctx context.Context, message *domain.ChatMessage) error {
	collection := r.db.Collection("messages")
	_, err := collection.InsertOne(ctx, bson.M{
		"sender":    message.Sender,
		"message":   message.Message,
		"timestamp": message.Timestamp,
	})
	return err
}
