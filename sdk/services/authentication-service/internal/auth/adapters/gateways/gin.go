package authgtw

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	mware "github.com/devpablocristo/golang/sdk/pkg/middleware/gin"
	sdk "github.com/devpablocristo/golang/sdk/pkg/rest/gin"
	sdkports "github.com/devpablocristo/golang/sdk/pkg/rest/gin/ports"
	dto "github.com/devpablocristo/golang/sdk/services/authentication-service/internal/auth/adapters/gateways/dto"

	ports "github.com/devpablocristo/golang/sdk/services/authentication-service/internal/auth/core/ports"
)

type GinHandler struct {
	ucs        ports.UseCases
	ginServer  sdkports.Server
	apiVersion string
	secret     string
}

func NewGinHandler(u ports.UseCases) *GinHandler {
	ginServer, err := sdk.Bootstrap()
	if err != nil {
		log.Fatalf("Gin Service error: %v", err)
	}

	return &GinHandler{
		ucs:        u,
		ginServer:  ginServer,
		apiVersion: ginServer.GetApiVersion(),
		secret:     viper.GetString("JWT_SECRET_KEY"),
	}
}

func (h *GinHandler) Start() error {
	h.routes()

	return h.ginServer.RunServer()
}

func (h *GinHandler) GetServer() sdkports.Server {
	return h.ginServer
}

func (h *GinHandler) routes() {
	router := h.ginServer.GetRouter()

	apiPrefix := "/api/" + h.apiVersion

	validated := router.Group(apiPrefix + "/auth/loginValidated")
	validated.Use(mware.ValidateLoginFields())
	{
		validated.POST("/login", h.Login)
	}

	authorized := router.Group(apiPrefix + "/auth/protected")
	authorized.Use(mware.JWTAuthMiddleware(h.secret))
	{
		authorized.GET("/auth-protected", h.ProtectedHandler)
	}
}

func (h *GinHandler) Login(c *gin.Context) {
	var req *mware.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	token, err := h.ucs.Login(c.Request.Context(), dto.LoginRequestToDomain(req))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, dto.LoginResponse{Token: token.AccessToken})
}

func (h *GinHandler) ProtectedHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok from protected"})
}
