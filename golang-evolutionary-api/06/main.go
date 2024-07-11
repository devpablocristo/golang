package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", home)
	router.GET("/hello", hello)
	router.POST("/bye", bye)

	router.Run(":8080")
}

func home(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to the home page!")
}

func hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello, world!")
}

func bye(c *gin.Context) {
	message := "Hello and bye!"
	c.String(http.StatusOK, "Received POST request with message: %s", message)
}
