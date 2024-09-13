package ports

import "github.com/golang-jwt/jwt/v5"

type JwtService interface {
	GenerateToken(claims jwt.MapClaims) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type JwtConfig interface {
	GetSecretKey() string
	Validate() error
}
