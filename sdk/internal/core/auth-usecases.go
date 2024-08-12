package core

import (
	"context"

	"github.com/devpablocristo/golang/sdk/cmd/gateways/auth/port"

	"github.com/devpablocristo/golang/sdk/internal/core/auth"
	"github.com/devpablocristo/golang/sdk/internal/core/user"
)

type AuthUseCasesPort interface {
	Login(context.Context, *user.User) (*auth.Auth, error)
}

type authUseCases struct {
	//userRepo  user.RepositoryPort
	//secretKey string
	broker port.MessageBroker
}

func NewAuthUseCases(b port.MessageBroker) AuthUseCasesPort {
	return &authUseCases{
		broker: b,
	}
}

// func (s *authUseCases) Login(ctx context.Context, username, password string) (string, error) {
// 	user, err := s.userRepo.GetUserByUsername(ctx, username)
// 	if err != nil {
// 		return "", errors.New("invalid credentials")
// 	}

// 	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
// 		// No revelar si la contraseña es incorrecta por seguridad.
// 		return "", errors.New("invalid credentials")
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"userID": user.UUID,
// 		"exp":    time.Now().Add(time.Hour * 72).Unix(),
// 	})

// 	// Firmar el token con la clave secreta.
// 	tokenString, err := token.SignedString([]byte(s.secretKey))
// 	if err != nil {
// 		return "", errors.New("could not sign token")
// 	}

// 	return tokenString, nil
// }

func (s *authUseCases) Login(ctx context.Context, user *user.User) (*auth.Auth, error) {
	// Enviar mensaje a RabbitMQ y obtener respuesta en los casos de uso
	// response, err := s.userRepo.GetUserByUsernameViaRabbitMQ(ctx, username)
	// if err != nil {
	// 	return "", errors.New("could not retrieve user information")
	// }

	// if err := bcrypt.CompareHashAndPassword([]byte(response.Password), []byte(password)); err != nil {
	// 	return "", errors.New("invalid credentials")
	// }

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"userID": response.UUID,
	// 	"exp":    time.Now().Add(time.Hour * 72).Unix(),
	// })

	// tokenString, err := token.SignedString([]byte(s.secretKey))
	// if err != nil {
	// 	return "", errors.New("could not sign token")
	// }

	return nil, nil
}
