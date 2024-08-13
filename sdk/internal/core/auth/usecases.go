package auth

import (
	"context"

	gtwports "github.com/devpablocristo/golang/sdk/cmd/gateways/auth/gtwports"

	"github.com/devpablocristo/golang/sdk/internal/core/auth/coreports"
	"github.com/devpablocristo/golang/sdk/internal/core/auth/entities"
)

type useAuthCases struct {
	broker gtwports.MessageBroker
}

func NewAuthUseCases(b gtwports.MessageBroker) coreports.AuthUseCases {
	return &useAuthCases{
		broker: b,
	}
}

func (s *useAuthCases) Login(ctx context.Context, user *entities.AuthUser) (*entities.Token, error) {
	userUUID, err := s.broker.GetUserUUID(ctx, user)
	if err != nil {
		return nil, err
	}

	_ = userUUID

	// Continuar con la lógica de autenticación
	token := &entities.Token{
		AccessToken: "generated-access-token",
		// Otras propiedades del token
	}

	return token, nil
}
