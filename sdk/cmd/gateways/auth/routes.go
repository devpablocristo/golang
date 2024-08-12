// package auth

// "github.com/gin-gonic/gin"

// wire "github.com/devpablocristo/golang/sdk/cmd/rest"
// is "github.com/devpablocristo/golang/sdk/pkg/init-setup"

// func AuthRoutes(r *gin.Engine) {
// 	authHandler, err := wire.InitializeAuthHandler()
// 	if err != nil {
// 		is.MicroLogError("authHandler error: %v", err)
// 	}

// 	api := r.Group("/api/v1")
// 	{
// 		api.POST("/login", authHandler.Login)
// 	}
// }

package auth

import (
	"github.com/gin-gonic/gin"

	mdhw "github.com/devpablocristo/golang/sdk/pkg/middleware"
)

func Routes(r *gin.Engine, ginHandler *GinHandler) {
	r.POST("/login", ginHandler.Login)

	secret := "secret"
	// "/api/v1/" <-- centralizar su creacion y enviarlo aqui
	authorized := r.Group("/api/v1/auth/protected")
	authorized.Use(mdhw.AuthMiddleware(secret))
	{
		authorized.GET("/auth-protected", ginHandler.ProtectedHandler)
	}
}
