package app

import (
	user "github.com/devpablocristo/interviews/bookstore-gin-rest-api/controllers/users"
)

func urlMap() {
	router.POST("/users", user.CreateUser)
	router.GET("/users", user.GetUsers)
	router.GET("/users/:id", user.GetUser)
	router.PUT("/users/:id", user.UpdateUser)
	router.DELETE("/users/:id", user.DeleteUser)
}
