package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/devpablocristo/qh/events/internal/core"
)

type AuthHandler struct {
	authUseCase core.AuthUseCasePort
}

func NewAuthHandler(authUseCase core.AuthUseCasePort) *AuthHandler {
	return &AuthHandler{
		authUseCase: authUseCase,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	token, err := h.authUseCase.Login(c.Request.Context(), req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, LoginResponse{Token: token})
}