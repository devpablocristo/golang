package api

import (
	"log"

	"github.com/gin-gonic/gin"

	hdl "github.com/devpablocristo/qh/analytics/cmd/api/handlers"
	ucs "github.com/devpablocristo/qh/analytics/internal/core"
)

func Build(dep *Dependencies) *gin.Engine {
	usecase := ucs.NewUseCase(dep.Repository)
	handler := hdl.NewRestHandler(usecase)

	r := gin.Default()

	v1 := r.Group("/api/v1/analytics")
	{
		v1.POST("/fake-create", handler.FakeCreateReport)
	}

	log.Println("Running server on port " + dep.RouterPort + "...")
	if err := r.Run(":" + dep.RouterPort); err != nil {
		log.Fatalf("Failed to run server: %s", err)
	}

	return r
}
