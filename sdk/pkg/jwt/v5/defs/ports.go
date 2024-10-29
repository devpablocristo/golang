package defs

// Config define la interfaz para la configuración del servicio JWT.
type Config interface {
	GetSecretKey() string
	Validate() error
}

// Service define la interfaz para el servicio JWT.
type Service interface {
	GenerateToken(claims Claims) (string, error)
	ValidateToken(tokenString string) (*TokenClaims, error)
}
