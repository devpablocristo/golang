package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/devpablocristo/golang/sdk/cmd/gateways/auth/dto"

	mdw "github.com/devpablocristo/golang/sdk/pkg/middleware/gin"

	"github.com/devpablocristo/golang/sdk/internal/core/auth/portscore"
)

type GinHandler struct {
	authUseCases portscore.AuthUseCases
}

func NewGinHandler(auc portscore.AuthUseCases) *GinHandler {
	return &GinHandler{
		authUseCases: auc,
	}
}

func (h *GinHandler) Login(c *gin.Context) {
	var req *mdw.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	token, err := h.authUseCases.Login(c.Request.Context(), dto.LoginRequestToDomain(req))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, dto.LoginResponse{Token: token.AccessToken})
}

func (h *GinHandler) ProtectedHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "ok from protected")
}
