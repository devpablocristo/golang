package authgtw

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"

	sdkmwr "github.com/devpablocristo/golang/sdk/pkg/middleware/gin"
	sdkgin "github.com/devpablocristo/golang/sdk/pkg/rest/gin"
	sdkginports "github.com/devpablocristo/golang/sdk/pkg/rest/gin/ports"
	sdktypes "github.com/devpablocristo/golang/sdk/pkg/types"
	sdkjwt "github.com/golang-jwt/jwt/v5"

	dto "github.com/devpablocristo/golang/sdk/ciudadanos/auth/internal/adapters/gateways/dto"
	ports "github.com/devpablocristo/golang/sdk/ciudadanos/auth/internal/core/ports"
)

type GinHandler struct {
	ucs        ports.UseCases
	ginServer  sdkginports.Server
	apiVersion string
	secret     string
}

func NewGinHandler(u ports.UseCases) (*GinHandler, error) {
	ginServer, err := sdkgin.Bootstrap()
	if err != nil {
		return nil, fmt.Errorf("Gin Service error: %w", err)
	}

	return &GinHandler{
		ucs:        u,
		ginServer:  ginServer,
		apiVersion: ginServer.GetApiVersion(),
		secret:     viper.GetString("JWT_SECRET_KEY"),
	}, nil
}

func (h *GinHandler) Start() error {
	h.routes()
	return h.ginServer.RunServer()
}

func (h *GinHandler) GetRouter() *gin.Engine {
	return h.ginServer.GetRouter()
}

func (h *GinHandler) routes() {
	router := h.ginServer.GetRouter()

	// Definir prefijos de ruta
	apiBase := "/api/" + h.apiVersion + "/auth"
	validatedPrefix := apiBase + "/validated"
	protectedPrefix := apiBase + "/protected"

	// Rutas públicas
	router.GET(apiBase+"/ping", h.Ping)

	// Rutas validadas (requieren validación de credenciales)
	validated := router.Group(validatedPrefix)
	{
		// Aplicar middleware de validación de credenciales
		validated.Use(sdkmwr.ValidateCredentials())
		validated.POST("/login", h.Login)
	}

	// Rutas protegidas (requieren JWT válido)
	authorized := router.Group(protectedPrefix)
	{
		// Aplicar middleware de validación JWT
		authorized.Use(sdkmwr.ValidateJwt(h.secret))
		authorized.GET("/protected-hi", h.ProtectedHi)
	}
}

func (h *GinHandler) AfipLogin(c *gin.Context) {
	cuit := c.Query("cuit")
	if cuit == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "CUIT is required"})
		return
	}

	err := h.ucs.AfipLogin(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "afip start error"})
		return
	}
}

func (h *GinHandler) Login(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")

	creds, exists := c.Get("creds:")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "credentials not found"})
		return
	}
	// Convertir el valor de `creds` al tipo correcto
	req, ok := creds.(sdktypes.LoginCredentials)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid credentials format"})
		return
	}

	token, err := h.ucs.Login(c.Request.Context(), dto.LoginRequestToDomain(&req))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, dto.LoginResponse{Token: token.AccessToken})
}

func (h *GinHandler) ProtectedHi(c *gin.Context) {
	token, exists := c.Get("token")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token not found"})
		return
	}

	// Realizar type assertion para extraer los claims
	claims, ok := token.(*sdkjwt.Token).Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token claims"})
		return
	}

	//Extraer los claims que sean necesarios (por ejemplo, el 'sub')
	sub := claims["sub"]

	c.JSON(http.StatusOK, gin.H{
		"message": []string{
			"hi! from protected.",
			"Value of the 'sub' claim from the token: " + sub.(string),
		},
	})

}

func (h *GinHandler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Pong!"})
}
