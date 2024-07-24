package userroutes

import (
	"github.com/gin-gonic/gin"

	wire "github.com/devpablocristo/qh/events/cmd/rest"
	is "github.com/devpablocristo/qh/events/pkg/init-setup"
	mdhw "github.com/devpablocristo/qh/events/pkg/middleware"
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
}
