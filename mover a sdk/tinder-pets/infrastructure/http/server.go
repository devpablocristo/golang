package http2

import (
	"github.com/gin-gonic/gin"
	application "github.com/luka385/tinder-pets/application/user_service"
)

func SetupsServer(UserUseCase application.UserServicePort) *gin.Engine {
	r := gin.Default()

	userHandler := NewUserHandler(UserUseCase)

	v1 := r.Group("/v1")
	{
		v1.POST("/users", userHandler.CreateUser)
		v1.PUT("/users/:id", userHandler.UpdateUser)
		v1.DELETE("/users/:id", userHandler.DeleteUser)
		v1.GET("/users/:id", userHandler.GetUserByID)
		v1.GET("/users", userHandler.GetAllUsers)
	}
	return r
}
