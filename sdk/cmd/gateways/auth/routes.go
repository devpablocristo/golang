package auth

import (
	mdw "github.com/devpablocristo/golang/sdk/pkg/middleware/gin"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, ginHandler *GinHandler, apiVersion string, secret string) {
	apiPrefix := "/api/" + apiVersion

	validated := r.Group(apiPrefix + "/auth/loginValidated")
	validated.Use(mdw.ValidateLoginFields())
	{
		validated.POST("/login", ginHandler.Login)
	}

	authorized := r.Group(apiPrefix + "/auth/protected")
	authorized.Use(mdw.JWTAuthMiddleware(secret))
	{
		authorized.GET("/auth-protected", ginHandler.ProtectedHandler)
	}
}
