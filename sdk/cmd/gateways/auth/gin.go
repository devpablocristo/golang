package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/devpablocristo/golang/sdk/cmd/gateways/auth/dto"
	ports "github.com/devpablocristo/golang/sdk/internal/core/auth/ports"
	mware "github.com/devpablocristo/golang/sdk/pkg/middleware/gin"
)

type GinHandler struct {
	ucs ports.AuthUseCases
}

func NewGinHandler(u ports.AuthUseCases) *GinHandler {
	return &GinHandler{
		ucs: u,
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
	c.JSON(http.StatusOK, "ok from protected")
}
