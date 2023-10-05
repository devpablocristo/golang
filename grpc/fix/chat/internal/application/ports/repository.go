package port

import (
	"context"

	"chat/internal/domain"
)

type Repository interface {
	SaveMessage(context.Context, *domain.ChatMessage) error
}
