package usergtw

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	sdkmwr "github.com/devpablocristo/golang/sdk/pkg/middleware/gin"
	sdkgin "github.com/devpablocristo/golang/sdk/pkg/rest/gin"
	sdkginports "github.com/devpablocristo/golang/sdk/pkg/rest/gin/ports"

	transport "github.com/devpablocristo/golang/sdk/sg/users/internal/adapters/gateways/transport"
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

	// Rutas para crear usuario
	// NOTE: pasarlo a protected
	router.POST(apiBase+"/create-user", h.CreateUser)

	// Rutas públicas
	router.GET(apiBase+"/ping", h.Ping)

	// Rutas validadas (requieren validación de credenciales)
	validated := router.Group(validatedPrefix)
	{
		// Aplicar middleware de validación de credenciales
		validated.Use(sdkmwr.ValidateCredentials())
		// Puedes añadir rutas aquí si es necesario
	}

	// Rutas protegidas (requieren JWT válido)
	protected := router.Group(protectedPrefix)
	{
		// Aplicar middleware de validación JWT
		protected.Use(sdkmwr.ValidateJwt(secrets["afip"]))
		protected.GET("/protected-hi", h.ProtectedHi)
	}
}

// AfipLogin maneja el login a través de AFIP (ejemplo)
func (h *GinHandler) AfipLogin(c *gin.Context) {
	cuil, err := afipJwtData(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JWT data: " + err.Error()})
		return
	}

	// Llamar a los use cases con el CUIL extraído del JWT
	// found, err := h.ucs.CheckUserStatus(c.Request.Context(), cuil)
	// if err != nil {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Error checking CUIL: " + err.Error()})
	// 	return
	// }

	// if found {
	// 	c.JSON(http.StatusOK, gin.H{"message": "CUIL found"})
	// 	return
	// }

	// Si todo es exitoso
	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "cuil": cuil})
}

// ProtectedHi responde desde una ruta protegida
func (h *GinHandler) ProtectedHi(c *gin.Context) {
	cuil, err := afipJwtData(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JWT data: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": []string{
			"Hi! From protected.",
			"CUIL: " + cuil,
		},
	})
}

// Ping responde con "Pong!" para verificar que el servidor está funcionando
func (h *GinHandler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Pong!"})
}

// CreateUser maneja la creación de un nuevo usuario
func (h *GinHandler) CreateUser(c *gin.Context) {
	var req *transport.CreateUserUser
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userUUID, err := h.ucs.CreateUser(c.Request.Context(), transport.ToUserDto(req))
	if err != nil {
		if err.Error() == "user already exists" {
			c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user: " + err.Error()})
		return
	}

	// Traducir la entidad de dominio a transport para la respuesta
	//usertransport := transport.ToUsertransport(user)

	c.JSON(http.StatusCreated, gin.H{
		"message": "user created",
		"uuid":    userUUID,
	})
}
