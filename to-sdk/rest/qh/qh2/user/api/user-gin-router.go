package api

import (
	"github.com/gin-gonic/gin"

	user "github.com/devpablocristo/golang/06-projects/qh/user/domain"
)

func GinRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/user", user.CreateUser)
	router.GET("/user", user.GetUsers)
	router.GET("/user/:id", user.GetUser)
	router.PUT("/user/:id", user.UpdateUser)
	router.DELETE("/user/:id", user.DeleteUser)
	return router
}
