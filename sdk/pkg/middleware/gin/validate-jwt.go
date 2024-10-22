package sdkmwr

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const (
	authHeaderName          = "Authorization"
	bearerPrefix            = "Bearer "
	errMissingAuthHeader    = "authorization header required"
	errInvalidSigningMethod = "unexpected signing method"
	errBearerPrefixRequired = "authorization header must start with Bearer"
	errInvalidToken         = "invalid token"
)

// ValidateJwt middleware para validar el JWT de forma agnóstica
func ValidateJwt(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtener el encabezado Authorization
		authHeader := c.GetHeader(authHeaderName)
		if authHeader == "" {
			abortWithError(c, http.StatusUnauthorized, errMissingAuthHeader)
			return
		}

		// Verificar que el encabezado empiece con "Bearer"
		if !strings.HasPrefix(authHeader, bearerPrefix) {
			abortWithError(c, http.StatusUnauthorized, errBearerPrefixRequired)
			return
		}

		// Extraer el token
		tokenString := strings.TrimPrefix(authHeader, bearerPrefix)

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
			abortWithError(c, http.StatusUnauthorized, errInvalidToken)
			return
		}

		// Guardar el token validado en el contexto para que los handlers lo utilicen
		c.Set("token", token)

		// Continuar con la siguiente función en la cadena de middlewares
		c.Next()
	}
}

// abortWithError centraliza la lógica para abortar con un error
func abortWithError(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{"error": message})
	c.Abort()
}
