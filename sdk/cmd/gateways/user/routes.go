package user

import (
	"github.com/gin-gonic/gin"

	mdw "github.com/devpablocristo/golang/sdk/pkg/middleware/gin"
)

func Routes(r *gin.Engine, h *GinHandler) {
	secret := "secret"
	authorized := r.Group("/api/v1/user/protected")
	authorized.Use(mdw.JWTAuthMiddleware(secret))
	{
		authorized.GET("/user-protected", h.CreateUser)
	}
}
