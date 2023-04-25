package main

import (
	"github.com/gin-gonic/gin"

	ginhandler "github.com/devpablocristo/golang/hex-arch/backend/internal/patients/infrastructure/handlers/gin"
)

const webServerPort string = ":8088"

func runPatientService(ph *ginhandler.GinHandler) {
	router := gin.New()
	setupPatientsRoutes(router, ph)
	router.Run(webServerPort)
}
