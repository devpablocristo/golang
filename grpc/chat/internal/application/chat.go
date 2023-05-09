package chat

import (
	"context"
	"time"

	"chat/internal/domain"
)

type Service struct {
	repo domain.ChatRepository
}

func NewService(repo domain.ChatRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) SendMessage(ctx context.Context, sender, message string) error {
	msg := &domain.ChatMessage{
		Sender:    sender,
		Message:   message,
		Timestamp: time.Now(),
	}
	return s.repo.SaveMessage(ctx, msg)
}
