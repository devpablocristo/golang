package auth

import (
	"context"
	"time"

	"github.com/devpablocristo/golang/sdk/cmd/gateways/auth/portsgtw"
	"github.com/devpablocristo/golang/sdk/internal/core/auth/entities"
	"github.com/devpablocristo/golang/sdk/internal/core/auth/portscore"
)

type useAuthCases struct {
	messageBroker portsgtw.MessageBroker
	accessControl portscore.AccessControl
}

func NewAuthUseCases(mb portsgtw.MessageBroker, ac portscore.AccessControl) portscore.AuthUseCases {
	return &useAuthCases{
		messageBroker: mb,
		accessControl: ac,
	}
}

func (s *useAuthCases) Login(ctx context.Context, credentials *entities.LoginCredentials) (*entities.Token, error) {
	// _, err := s.messageBroker.GetUserUUID(ctx, credentials)
	// if err != nil {
	// 	return nil, err
	// }

	return &entities.Token{
		AccessToken: "",
		ExpiresAt:   time.Now().Add(time.Hour * 24),
	}, nil
}
