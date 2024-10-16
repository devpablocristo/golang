package authgtw

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	mware "github.com/devpablocristo/golang/sdk/pkg/middleware/gin"
	sdk "github.com/devpablocristo/golang/sdk/pkg/rest/gin"
	sdkports "github.com/devpablocristo/golang/sdk/pkg/rest/gin/ports"
	dto "github.com/devpablocristo/golang/sdk/qh/rating/internal/adapters/gateways/dto"
	ports "github.com/devpablocristo/golang/sdk/qh/rating/internal/core/ports"
)

type GinHandler struct {
	ucs        ports.RatUseCases
	ginServer  sdkports.Server
	apiVersion string
}

func NewGinHandler(u ports.RatUseCases) (*GinHandler, error) {
	ginServer, err := sdk.Bootstrap()
	if err != nil {
		return nil, fmt.Errorf("Gin Service error: %w", err)
	}

	return &GinHandler{
		ucs:        u,
		ginServer:  ginServer,
		apiVersion: ginServer.GetApiVersion(),
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

	validated := router.Group(apiPrefix + "/rating/loginValidated")
	validated.Use(mware.ValidateCredentials())
	{
		validated.POST("/login", h.Login)
	}

	authorized := router.Group(apiPrefix + "/rating/protected")
	authorized.Use(mware.ValidateJwt(h.secret))
	{
		authorized.GET("/rating-protected", h.ProtectedHandler)
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
