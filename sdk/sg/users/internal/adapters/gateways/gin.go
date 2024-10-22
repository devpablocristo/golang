package usergtw

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	sdkmwr "github.com/devpablocristo/golang/sdk/pkg/middleware/gin"
	sdkgin "github.com/devpablocristo/golang/sdk/pkg/rest/gin"
	sdkginports "github.com/devpablocristo/golang/sdk/pkg/rest/gin/ports"

	ports "github.com/devpablocristo/golang/sdk/sg/users/internal/core/ports"
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
	apiBase := "/api/" + apiVersion + "/users"
	validatedPrefix := apiBase + "/validated"
	protectedPrefix := apiBase + "/protected"

	// Rutas públicas
	router.GET(apiBase+"/ping", h.Ping)

	// Rutas validadas (requieren validación de credenciales)
	validated := router.Group(validatedPrefix)
	{
		// Aplicar middleware de validación de credenciales
		validated.Use(sdkmwr.ValidateCredentials())
	}

	// Rutas protegidas (requieren JWT válido)
	afipusersorized := router.Group(protectedPrefix)
	{
		// Aplicar middleware de validación JWT
		afipusersorized.Use(sdkmwr.ValidateJwt(secrets["afip"]))
		afipusersorized.GET("/protected-hi", h.ProtectedHi)
	}
}

func (h *GinHandler) AfipLogin(c *gin.Context) {
	cuit, err := afipJwtData(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid jwt data: " + err.Error()})
		return
	}

	// Llamar a los use cases con el CUIT extraído del JWT
	found, err := h.ucs.CheckCuit(c.Request.Context(), cuit)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "error checking cuit: " + err.Error()})
		return
	}

	if found {
		c.JSON(http.StatusOK, gin.H{"message": "cuit founded"})
		return
	}

	// Si todo es exitoso
	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "cuit": cuit})
}

func (h *GinHandler) ProtectedHi(c *gin.Context) {
	cuit, err := afipJwtData(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid jwt data: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": []string{
			"hi! from protected.",
			"cuit: " + cuit,
		},
	})

}

func (h *GinHandler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Pong!"})
}
