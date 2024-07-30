package validate

import (
	"context"
	"errors"

	port "github.com/devpablocristo/99minutos/order-manager/internal/application/port"
)

type ValidateUser struct {
	userRepository port.UserRepo
}

func NewValidateUser(ur port.UserRepo) port.ValidateUser {
	return &ValidateUser{
		userRepository: ur,
	}
}

func (uc *ValidateUser) Execute(ctx context.Context, email, password string) (int16, error) {
	user, err := uc.userRepository.FindByEmail(ctx, email)
	if err != nil {
		return 0, err
	}

	if user.Password != password {
		return 0, errors.New("invalid password")
	}

	return user.Role, nil
}
