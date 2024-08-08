package core

import (
	"context"

	"github.com/devpablocristo/golang/sdk/internal/core/user"
)

type GreaterUseCasesPort interface {
	Login(context.Context, string, string) (string, error)
}

type greeterUseCases struct {
	userRepo  user.RepositoryPort
	secretKey string
}

func NewGreaterUseCases(ur user.RepositoryPort, sk string) GreaterUseCasesPort {
	return &authUseCases{
		userRepo:  ur,
		secretKey: sk,
	}
}

func (s *authUseCases) Hello(ctx context.Context) (string, error) {
	return "hello", nil
}
