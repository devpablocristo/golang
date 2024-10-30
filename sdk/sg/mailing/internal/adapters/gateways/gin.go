package mailgtw

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	sdkmwr "github.com/devpablocristo/golang/sdk/pkg/middleware/gin"
	sdkgin "github.com/devpablocristo/golang/sdk/pkg/rest/gin"
	sdkgindefs "github.com/devpablocristo/golang/sdk/pkg/rest/gin/defs"

	ports "github.com/devpablocristo/golang/sdk/sg/mailing/internal/core/ports"
)

type GinHandler struct {
	ucs       ports.UseCases
	ginServer sdkgindefs.Server
}

func NewGinHandler(u ports.UseCases) (*GinHandler, error) {
	// Aquí aceptará el tipo ports.UseCases
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
	}

	// Rutas protegidas (requieren JWT válido)
	afipAuthorized := router.Group(protectedPrefix)
	{
		// Aplicar middleware de validación JWT
		afipAuthorized.Use(sdkmwr.ValidateJwt(secrets["afip"]))
		afipAuthorized.GET("/protected-hi", h.ProtectedHi)
	}
}

func (h *GinHandler) ProtectedHi(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": []string{
			"hi! from protected.",
		},
	})

}

func (h *GinHandler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Pong!"})
}
