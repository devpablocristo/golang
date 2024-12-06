package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", homeGet)
	router.POST("/", homePost)
	router.PUT("/", homePut)
	router.DELETE("/", homeDelete)

	router.GET("/hello", hello)

	router.POST("/bye", bye)

	router.Run(":8080")
}

func homeGet(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to the home page!")
}

func homePost(c *gin.Context) {
	c.String(http.StatusOK, "Post to the home page!")
}

func homePut(c *gin.Context) {
	c.String(http.StatusOK, "Put to the home page!")
}

func homeDelete(c *gin.Context) {
	c.String(http.StatusOK, "Delete the home page!")
}

func hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello, world!")
}

func bye(c *gin.Context) {
	c.String(http.StatusOK, "Goodbye guys!")
}
