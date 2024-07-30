package app

// router solo es accesible desde app.go

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

var (
	router     = gin.Default()
	serverPort = "8001"
)

func StartApp() {

	urlMap()
	err := router.Run(":" + serverPort)

	logErrors(err)
}

func logErrors(err error) {
	if err != nil {
		log.Fatal("ERROR!!! ", err)
	}
}
