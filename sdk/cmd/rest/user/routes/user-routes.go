package userroutes

import (
	"github.com/gin-gonic/gin"

	wire "github.com/devpablocristo/golang-sdk/cmd/rest"
	is "github.com/devpablocristo/golang-sdk/pkg/init-setup"
	mdhw "github.com/devpablocristo/golang-sdk/pkg/middleware"
)

func UserRoutes(r *gin.Engine) {
	userHandler, err := wire.InitializeUserHandler()
	if err != nil {
		is.MicroLogError("userHandler error: %v", err)
	}

	secret := "secret"
	user := r.Group("/api/v1/user")
	user.Use(mdhw.AuthMiddleware(secret))
	{
		user.GET(":id", userHandler.GetUser)
	}

	// Ruta de Salud
	r.GET("/health", userHandler.Health)
	r.GET("/ping", userHandler.Ping)
}
