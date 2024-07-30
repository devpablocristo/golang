package api

import (
	"log"

	"github.com/gin-gonic/gin"

	hdl "github.com/devpablocristo/dive-challenge/cmd/api/handlers"
	ucs "github.com/devpablocristo/dive-challenge/internal/core"
)

func Build(dep *Dependencies) *gin.Engine {
	usecase := ucs.NewUseCase(dep.Repository, dep.ApiClient)
	handler := hdl.NewRestHandler(usecase)

	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/ltp", handler.GetLTP)
	}

	log.Println("Running server on port " + dep.RouterPort + "...")
	if err := r.Run(":" + dep.RouterPort); err != nil {
		log.Fatalf("Failed to run server: %s", err)
	}

	return r
}
