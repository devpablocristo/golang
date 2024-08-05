package user

import (
	"github.com/gin-gonic/gin"

	wire "github.com/devpablocristo/golang/sdk/cmd/rest"
	is "github.com/devpablocristo/golang/sdk/pkg/init-setup"
	mdhw "github.com/devpablocristo/golang/sdk/pkg/middleware"
)

func Routes(r *gin.Engine) {
	handler, err := wire.InitializeUserHandler()
	if err != nil {
		is.MicroLogError("userHandler error: %v", err)
	}

	secret := "secret"
	user := r.Group("/api/v1/user")
	user.Use(mdhw.AuthMiddleware(secret))
	{
		user.GET(":id", handler.GetUser)
	}

}
