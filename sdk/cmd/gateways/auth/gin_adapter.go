package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/devpablocristo/golang/sdk/internal/core"
)

type GinHandler struct {
	authUseCases core.AuthUseCasesPort
}

func NewGinHandler(authUseCases core.AuthUseCasesPort) *GinHandler {
	return &GinHandler{
		authUseCases: authUseCases,
	}
}

func (h *GinHandler) Login(c *gin.Context) {
	var req *LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	auth, err := h.authUseCases.Login(c.Request.Context(), toDomain(req))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, LoginResponse{Token: auth.Token})
}

func (h *GinHandler) ProtectedHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "ok from protected")
}
