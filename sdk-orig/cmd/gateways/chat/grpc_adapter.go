package chat

import (
	"io"
	"log"

	"github.com/devpablocristo/golang/sdk/cmd/gateways/chat/pb"
	"github.com/devpablocristo/golang/sdk/internal/core"
)

type ChatHandler struct {
	chatService core.ChatUseCasesPort
	pb.ChatServiceServer
}

func NewChatHandler(cs core.ChatUseCasesPort) *ChatHandler {
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
