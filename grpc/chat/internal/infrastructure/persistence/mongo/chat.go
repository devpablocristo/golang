package mongo

import (
	"context"

	"chat-api/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ChatRepository struct {
	db *mongo.Database
}

func NewChatRepository(db *mongo.Database) *ChatRepository {
	return &ChatRepository{db: db}
}

func (r *ChatRepository) SaveMessage(ctx context.Context, message *domain.ChatMessage) error {
	collection := r.db.Collection("messages")
	_, err := collection.InsertOne(ctx, bson.M{
		"sender":    message.Sender,
		"message":   message.Message,
		"timestamp": message.Timestamp,
	})
	return err
}
