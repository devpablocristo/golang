package chatservice

import (
	"context"
	"time"

	port "chat/internal/application/ports"
	domain "chat/internal/domain"
	pb "chat/internal/pb"
)

type ChatService struct {
	repo port.Repository
	pb.ChatServiceServer
}

func NewChatService(repo port.Repository) port.ChatService {
	return &ChatService{
		repo: repo,
	}
}

func (s *ChatService) SendMessage(ctx context.Context, sender, message string) error {
	msg := &domain.ChatMessage{
		Sender:    sender,
		Message:   message,
		Timestamp: time.Now(),
	}
	return s.repo.SaveMessage(ctx, msg)
}
