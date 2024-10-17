package sdkmwr

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const (
	errMissingAuthHeader    = "authorization header required"
	errInvalidSigningMethod = "unexpected signing method"
	errBearerPrefixRequired = "authorization header must start with Bearer"
	errInvalidToken         = "invalid token"
)

// ValidateJwt middleware para validar el JWT de forma agnóstica
func ValidateJwt(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtener el encabezado Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": errMissingAuthHeader})
			c.Abort()
			return
		}

		// Verificar que el encabezado empiece con "Bearer"
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": errBearerPrefixRequired})
			c.Abort()
			return
		}

		// Extraer el token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Validar el token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
			// Verificar que el método de firma sea HMAC
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New(errInvalidSigningMethod)
			}
			// Retornar la clave secreta para validar la firma
			return []byte(secretKey), nil
		})

		// Si el token no es válido, devolver error
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": errInvalidToken})
			c.Abort()
			return
		}

		// Guardar el token validado en el contexto para que los handlers lo utilicen
		c.Set("token", token)

		// Continuar con la siguiente función en la cadena de middlewares
		c.Next()
	}
}
