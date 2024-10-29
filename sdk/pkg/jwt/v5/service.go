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

// NewService inicializa el servicio JWT con la configuración proporcionada.
func newService(c defs.Config) (defs.Service, error) {
	if err := c.Validate(); err != nil {
		return nil, err
	}
	return &service{
		secretKey: []byte(c.GetSecretKey()),
	}, nil
}

func (s *service) GenerateToken(claims defs.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(s.secretKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}
	return signedToken, nil
}

func (j *service) ValidateToken(tokenString string) (*defs.TokenClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de firma inesperado: %v", token.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("error al validar el token: %w", err)
	}

	// Verifica que el token sea válido y extrae las claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("token inválido")
	}

	// Extrae las claims necesarias
	subject, ok := claims["sub"].(string)
	if !ok {
		return nil, fmt.Errorf("el token no contiene 'sub'")
	}

	expFloat, ok := claims["exp"].(float64)
	if !ok {
		return nil, fmt.Errorf("el token no contiene 'exp'")
	}
	expiresAt := time.Unix(int64(expFloat), 0)

	iatFloat, ok := claims["iat"].(float64)
	if !ok {
		return nil, fmt.Errorf("el token no contiene 'iat'")
	}
	issuedAt := time.Unix(int64(iatFloat), 0)

	// Crea una estructura TokenClaims con la información extraída
	tokenClaims := &defs.TokenClaims{
		Subject:   subject,
		ExpiresAt: expiresAt,
		IssuedAt:  issuedAt,
	}

	return tokenClaims, nil
}
