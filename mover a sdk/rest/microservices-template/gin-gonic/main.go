package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {

	r := setupRouter()
	containerPort := ":4000"

	// listen and serve on 0.0.0.0:8080
	//r.Run()

	err := r.Run(containerPort)
	logErrors(err)
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hola mundo!")
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "super nana",
		})
	})

	return r
}

func logErrors(err error) {
	if err != nil {
		log.Fatal("ERROR!!! ", err)
	}
}
