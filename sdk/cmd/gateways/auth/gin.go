package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/devpablocristo/golang/sdk/cmd/gateways/auth/dto"
	"github.com/devpablocristo/golang/sdk/internal/core/auth/coreports"
)

type GinHandler struct {
	useCases coreports.AuthUseCases
}

func NewGinHandler(useCases coreports.AuthUseCases) *GinHandler {
	return &GinHandler{
		useCases: useCases,
	}
}

func (h *GinHandler) Login(c *gin.Context) {
	var req *dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	token, err := h.useCases.Login(c.Request.Context(), dto.LoginRequestToDomain(req))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, dto.LoginResponse{Token: token.AccessToken})
}

func (h *GinHandler) ProtectedHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "ok from protected")
}
