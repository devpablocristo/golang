package port

import "context"

type ChatService interface {
	SendMessage(context.Context, string, string) error
}
