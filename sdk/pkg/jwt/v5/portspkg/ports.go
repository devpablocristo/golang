package portspkg

import "github.com/golang-jwt/jwt/v5"

// JWTService define las operaciones básicas para manejar tokens JWT.
type JWTClient interface {
	GenerateToken(claims jwt.MapClaims) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}
