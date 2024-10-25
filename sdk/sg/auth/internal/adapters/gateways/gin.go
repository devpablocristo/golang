package authgtw

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	sdkmwr "github.com/devpablocristo/golang/sdk/pkg/middleware/gin"
	sdkgin "github.com/devpablocristo/golang/sdk/pkg/rest/gin"
	sdkginports "github.com/devpablocristo/golang/sdk/pkg/rest/gin/ports"
	sdktypes "github.com/devpablocristo/golang/sdk/pkg/types"

	dto "github.com/devpablocristo/golang/sdk/sg/auth/internal/adapters/gateways/dto"
	ports "github.com/devpablocristo/golang/sdk/sg/auth/internal/core/ports"
)

type GinHandler struct {
	ucs       ports.UseCases
	ginServer sdkginports.Server
}

func NewGinHandler(u ports.UseCases) (*GinHandler, error) {
	ginServer, err := sdkgin.Bootstrap()
	if err != nil {
		return nil, fmt.Errorf("Gin Service error: %w", err)
	}

	return &GinHandler{
		ucs:       u,
		ginServer: ginServer,
	}, nil
}

func (h *GinHandler) GetRouter() *gin.Engine {
	return h.ginServer.GetRouter()
}

func (h *GinHandler) Start() error {
	// Cargar el secret solo cuando sea necesario
	secrets, err := getSecrets()
	if err != nil {
		return fmt.Errorf("failed to load secrets: %w", err)
	}

	// Configurar rutas
	h.routes(secrets)

	// Iniciar el servidor
	return h.ginServer.RunServer()
}

func (h *GinHandler) routes(secrets map[string]string) {
	router := h.ginServer.GetRouter()

	// Definir prefijos de ruta
	apiVersion := h.ginServer.GetApiVersion()
	apiBase := "/api/" + apiVersion + "/auth"
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
	afipAuthorized := router.Group(protectedPrefix)
	{
		// Aplicar middleware de validación JWT
		afipAuthorized.Use(sdkmwr.ValidateJwt(secrets["afip"]))
		afipAuthorized.GET("/protected-hi", h.ProtectedHi)
	}
}

func (h *GinHandler) AfipLogin(c *gin.Context) {
	cuil, err := afipJwtData(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid jwt data: " + err.Error()})
		return
	}

	// Llamar a los use cases con el CUIL extraído del JWT
	if err := h.ucs.AfipLogin(c.Request.Context(), cuil); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "afip start error"})
		return
	}

	// Si todo es exitoso
	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "cuil": cuil})
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
	cuil, err := afipJwtData(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid jwt data: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": []string{
			"hi! from protected.",
			"cuil: " + cuil,
		},
	})

}

func (h *GinHandler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Pong!"})
}
