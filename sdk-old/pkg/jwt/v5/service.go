// sdkjwt/service.go

package sdkjwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/devpablocristo/golang/sdk/pkg/jwt/v5/defs"
)

type service struct {
	secretKey []byte
}

func newService(c defs.Config) (defs.Service, error) {
	if err := c.Validate(); err != nil {
		return nil, err
	}
	return &service{
		secretKey: []byte(c.GetSecretKey()),
	}, nil
}

func (s *service) GenerateTokenForSubject(subject string, expiration time.Duration) (string, error) {
	claims := defs.Claims{
		Subject: subject,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(s.secretKey)
	if err != nil {
		return "", fmt.Errorf("error al firmar el token: %w", err)
	}
	return signedToken, nil
}

func (s *service) ValidateToken(tokenString string) (*defs.TokenClaims, error) {
	claims := &defs.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de firma inesperado: %v", token.Header["alg"])
		}
		return s.secretKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("error al validar el token: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("token inválido")
	}

	tokenClaims := &defs.TokenClaims{
		Subject:   claims.Subject,
		ExpiresAt: claims.ExpiresAt.Time,
		IssuedAt:  claims.IssuedAt.Time,
	}

	return tokenClaims, nil
}
