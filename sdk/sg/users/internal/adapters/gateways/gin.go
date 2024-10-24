package usergtw

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	sdkmwr "github.com/devpablocristo/golang/sdk/pkg/middleware/gin"
	sdkgin "github.com/devpablocristo/golang/sdk/pkg/rest/gin"
	sdkginports "github.com/devpablocristo/golang/sdk/pkg/rest/gin/ports"

	"github.com/devpablocristo/golang/sdk/sg/users/internal/adapters/gateways/dto"
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

	// NOTE: mover a protected
	router.POST(apiBase+"/create-user", h.CreateUser)

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
	found, err := h.ucs.CheckUserStatus(c.Request.Context(), cuit)
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

func (h *GinHandler) CreateUser(c *gin.Context) {
	// NOTE: cambiar a DTO
	var req *dto.UserDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validar que solo uno de los UUIDs esté presente
	if (req.PersonUUID != nil && req.CompanyUUID != nil) || (req.PersonUUID == nil && req.CompanyUUID == nil) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Debe proporcionar solo uno de person_uuid o company_uuid"})
		return
	}

	// Llamar al caso de uso para crear el usuario
	if err := h.ucs.CreateUser(c.Request.Context(), dto.FromUserDTO(req)); err != nil {
		if err.Error() == "el usuario ya existe" {
			c.JSON(http.StatusConflict, gin.H{"error": "El usuario ya existe"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el usuario: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})
}

// hashPassword genera un hash seguro para la contraseña
func hashPassword(password string) (string, error) {
	// Implementa el hashing de la contraseña, por ejemplo usando bcrypt
	// Este es un ejemplo utilizando bcrypt
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}
