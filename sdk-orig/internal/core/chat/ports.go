package chat

import (
	"context"
)

type Repository interface {
	SaveMessage(context.Context, *ChatMessage) error
}
