package auth

import (
	"github.com/gin-gonic/gin"

	mware "github.com/devpablocristo/golang/sdk/pkg/middleware/gin"
)

func Routes(r *gin.Engine, ginHandler *GinHandler, apiVersion string, secret string) {
	apiPrefix := "/api/" + apiVersion

	validated := r.Group(apiPrefix + "/auth/loginValidated")
	validated.Use(mware.ValidateLoginFields())
	{
		validated.POST("/login", ginHandler.Login)
	}

	authorized := r.Group(apiPrefix + "/auth/protected")
	authorized.Use(mware.JWTAuthMiddleware(secret))
	{
		authorized.GET("/auth-protected", ginHandler.ProtectedHandler)
	}
}
