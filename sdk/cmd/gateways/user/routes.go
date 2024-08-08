package user

import (
	"github.com/gin-gonic/gin"

	mdhw "github.com/devpablocristo/golang/sdk/pkg/middleware"
)

// TODO: resolver como usar mejor routes, no me gusta tener q hacer una paquete diferente
func Routes(r *gin.Engine, handler *Handler) {

	secret := "secret"
	user := r.Group("/api/v1/user")
	user.Use(mdhw.AuthMiddleware(secret))
	{
		user.GET(":id", handler.GetUser)
	}
}
