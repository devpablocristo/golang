package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/", helloWorld)

	log.Println("Server started at http://localhost:8080/")

	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

func helloWorld(c *gin.Context) {
	c.String(http.StatusOK, "¡Hello World!")
}
