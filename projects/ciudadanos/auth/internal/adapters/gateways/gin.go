package authgtw

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	dto "github.com/devpablocristo/golang/sdk/ciudadanos/auth/internal/adapters/gateways/dto"
	ports "github.com/devpablocristo/golang/sdk/ciudadanos/auth/internal/core/ports"
	mware "github.com/devpablocristo/golang/sdk/middleware/gin"
	sdk "github.com/devpablocristo/golang/sdk/rest/gin"
	sdkports "github.com/devpablocristo/golang/sdk/rest/gin/ports"
)

type GinHandler struct {
	ucs        ports.UseCases
	ginServer  sdkports.Server
	apiVersion string
	secret     string
}

func NewGinHandler(u ports.UseCases) (*GinHandler, error) {
	ginServer, err := sdk.Bootstrap()
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

	apiPrefix := "/api/" + h.apiVersion

	router.GET(apiPrefix+"/ping", h.Ping)

	validated := router.Group(apiPrefix + "/authe/loginValidated")
	validated.Use(mware.ValidateCredentials())
	{
		validated.POST("/login", h.Login)
	}

	authorized := router.Group(apiPrefix + "/authe/protected")
	authorized.Use(mware.ValidateJwt(h.secret))
	{
		authorized.GET("/authe-protected", h.ProtectedHandler)
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

func (h *GinHandler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Pong!"})
}
