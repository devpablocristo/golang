package grpc

import (
	"io"
	"log"

	"chat-api/internal/application/chat"
	"chat-api/pkg/chat"
)

type ChatServer struct {
	service chat.UnimplementedChatServiceServer
}

func NewChatServer(service *chat.Service) *ChatServer {
	return &ChatServer{
		service: service,
	}
}

func (s *ChatServer) StreamChat(stream chat.ChatService_StreamChatServer) error {
	for {
		incomingMsg, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Printf("Error receiving message from stream: %v", err)
			return err
		}

		err = s.service.SendMessage(stream.Context(), incomingMsg.Sender, incomingMsg.Message)
		if err != nil {
			log.Printf("Error saving message: %v", err)
			return err
		}

		if err := stream.Send(incomingMsg); err != nil {
			log.Printf("Error sending message to stream: %v", err)
			return err
		}
	}
}
