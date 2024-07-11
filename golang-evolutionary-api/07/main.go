package main

import (
	"io"
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
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}
	message := string(body)
	c.String(http.StatusOK, "Received POST request with message: %s", message)
}
