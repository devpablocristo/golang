package domain

import (
	"context"
	"time"
)

type ChatMessage struct {
	Sender    string
	Message   string
	Timestamp time.Time
}

type ChatRepository interface {
	SaveMessage(ctx context.Context, message *ChatMessage) error
}
