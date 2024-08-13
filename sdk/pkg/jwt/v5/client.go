package jwtoken

import (
	"fmt"
	"sync"

	"github.com/golang-jwt/jwt/v5"

	"github.com/devpablocristo/golang/sdk/pkg/jwt/v5/pkgports"
)

var (
	jwtInstance pkgports.JWTClient
	jwtOnce     sync.Once
	errInit     error
)

type jwtService struct {
	secretKey string
}

func InitializeJWTService(secretKey string) error {
	jwtOnce.Do(func() {
		if secretKey == "" {
			errInit = fmt.Errorf("secret key cannot be empty")
			return
		}

		jwtInstance = &jwtService{
			secretKey: secretKey,
		}
	})
	return errInit
}

// GetJWTInstance devuelve la instancia del servicio JWT.
func GetJWTInstance() (pkgports.JWTClient, error) {
	if jwtInstance == nil {
		return nil, fmt.Errorf("JWT service is not initialized")
	}
	return jwtInstance, nil
}

// GenerateToken genera un token JWT con las declaraciones proporcionadas.
func (j *jwtService) GenerateToken(claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// ValidateToken valida un token JWT y devuelve el token decodificado si es válido.
func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Verifica que el método de firma coincide con el esperado
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
}
