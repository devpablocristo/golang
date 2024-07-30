package app

// router solo es accesible desde app.go

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()

}

var (
	router = gin.Default()
)

func StartApp() {

	serverPort := os.Getenv("SERV_LOCL_PORT")

	urlMap()
	fmt.Println(serverPort)
	err := router.Run(":" + serverPort)

	logErrors(err)
}

func logErrors(err error) {
	if err != nil {
		log.Fatal("ERROR!!! ", err)
	}
}
