package grpchandler

import (
	"io"
	"log"

	port "chat/internal/application/ports"
	pb "chat/internal/pb"
)

type ChatHandler struct {
	chatService port.ChatService
	pb.ChatServiceServer
}

func NewChatHandler(cs port.ChatService) *ChatHandler {
	return &ChatHandler{
		chatService: cs,
	}
}

func (s *ChatHandler) StreamChat(stream pb.ChatService_ChatServer) error {
	for {
		incomingMsg, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Printf("Error receiving message from stream: %v", err)
			return err
		}

		err = s.chatService.SendMessage(stream.Context(), incomingMsg.Sender, incomingMsg.Message)
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
