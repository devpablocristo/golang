package jwtpkg

import (
	"fmt"
	"sync"

	"github.com/devpablocristo/golang/sdk/pkg/jwt/v5/portspkg"
	"github.com/golang-jwt/jwt/v5"
)

var (
	jwtInstance portspkg.JwtClient
	jwtOnce     sync.Once
	initError   error
)

type jwtService struct {
	secretKey string
}

func InitializeJWTService(secretKey string) error {
	jwtOnce.Do(func() {
		if secretKey == "" {
			initError = fmt.Errorf("secret key cannot be empty")
			return
		}

		jwtInstance = &jwtService{
			secretKey: secretKey,
		}
	})
	return initError
}

func GetJWTInstance() (portspkg.JwtClient, error) {
	if jwtInstance == nil {
		return nil, fmt.Errorf("JWT service is not initialized")
	}
	return jwtInstance, nil
}

func (j *jwtService) GenerateToken(claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
}