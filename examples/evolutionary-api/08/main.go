package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", home)
	router.GET("/hello", hello)
	router.POST("/bye", bye)

	fmt.Println("Welcome!")
	log.Println("Server started at http://localhost:8080/")

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func home(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to the home page!")
}

func hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello, world!")
}

func bye(c *gin.Context) {
	var msg map[string]string
	if err := c.BindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind JSON"})
		return
	}
	message, exists := msg["message"]
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Message field is missing"})
		return
	}
	c.String(http.StatusOK, "Received POST request with message: %s", message)
}
