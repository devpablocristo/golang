package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"

	dto "github.com/devpablocristo/golang/sdk/cmd/gateways/auth/dto"
	ports "github.com/devpablocristo/golang/sdk/internal/core/auth/ports"
	mware "github.com/devpablocristo/golang/sdk/pkg/middleware/gin"
	sdkgin "github.com/devpablocristo/golang/sdk/pkg/rest/gin/ports"
)

type GinHandler struct {
	ucs        ports.AuthUseCases
	ginServer  sdkgin.Server
	apiVersion string
	secret     string
}

func NewGinHandler(u ports.AuthUseCases, ginServer sdkgin.Server, apiVersion string, secret string) *GinHandler {
	return &GinHandler{
		ucs:        u,
		ginServer:  ginServer,
		apiVersion: apiVersion,
		secret:     secret,
	}
}

func (h *GinHandler) Start() error {
	h.routes()

	return h.ginServer.RunServer()
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

// ProtectedHandler maneja el acceso a rutas protegidas.
func (h *GinHandler) ProtectedHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok from protected"})
}
