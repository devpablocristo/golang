package core

import (
	"context"
	"errors"
	"time"

	"github.com/devpablocristo/qh/events/internal/core/user"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase struct {
	userRepo  user.RepositoryPort
	secretKey string
}

func NewAuthUseCase(ur user.RepositoryPort, sk string) AuthUseCasePort {
	return &AuthUseCase{
		userRepo:  ur,
		secretKey: sk,
	}
}

func (s *AuthUseCase) Login(ctx context.Context, username, password string) (string, error) {
	user, err := s.userRepo.GetUserByUsername(ctx, username)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		// No revelar si la contraseña es incorrecta por seguridad.
		return "", errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": user.ID,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	})

	// Firmar el token con la clave secreta.
	tokenString, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", errors.New("could not sign token")
	}

	return tokenString, nil
}