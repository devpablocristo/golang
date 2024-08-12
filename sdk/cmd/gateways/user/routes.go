package user

import (
	"github.com/gin-gonic/gin"

	mdhw "github.com/devpablocristo/golang/sdk/pkg/middleware"
)

func Routes(r *gin.Engine, ginHandler *GinHandler) {
	secret := "secret"
	// "/api/v1/" <-- centralizar su creacion y enviarlo aqui
	authorized := r.Group("/api/v1/user/protected")
	authorized.Use(mdhw.AuthMiddleware(secret))
	{

		authorized.GET("/user-protected", ginHandler.CreateUser)
	}
}
