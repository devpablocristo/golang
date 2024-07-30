package authroutes

import (
	"github.com/gin-gonic/gin"

	wire "github.com/devpablocristo/qh/events/cmd/rest"
	is "github.com/devpablocristo/qh/events/pkg/init-setup"
)

func AuthRoutes(r *gin.Engine) {
	authHandler, err := wire.InitializeAuthHandler()
	if err != nil {
		is.MicroLogError("authHandler error: %v", err)
	}

	api := r.Group("/api/v1")
	{
		api.POST("/login", authHandler.Login)
	}
}
