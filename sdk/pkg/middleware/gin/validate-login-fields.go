package sdkmiddleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username     string `json:"username" binding:"required"`
	PasswordHash string `json:"password" binding:"required"`
}

func ValidateLoginFields() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var loginRequest LoginRequest
		if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Username and password are required"})
			ctx.Abort()
			return
		}

		if loginRequest.Username == "" || loginRequest.PasswordHash == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Username and password cannot be empty"})
			ctx.Abort()
			return
		}

		ctx.Set("loginRequest", loginRequest)
		ctx.Next()
	}
}
