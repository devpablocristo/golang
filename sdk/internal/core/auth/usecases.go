package auth

import (
	"context"
	"time"

	entities "github.com/devpablocristo/golang/sdk/internal/core/auth/entities"
	ports "github.com/devpablocristo/golang/sdk/internal/core/auth/ports"
)

type useCases struct {
	accessControl ports.AccessControl
}

func NewUseCases(ac ports.AccessControl) ports.UseCases {
	return &useCases{
		accessControl: ac,
	}
}

func (s *useCases) Login(ctx context.Context, credentials *entities.LoginCredentials) (*entities.Token, error) {
	// _, err := s.messageBroker.GetUserUUID(ctx, credentials)
	// if err != nil {
	// 	return nil, err
	// }

	return &entities.Token{
		AccessToken: "",
		ExpiresAt:   time.Now().Add(time.Hour * 24),
	}, nil
}
