package main

import (
	"github.com/gin-gonic/gin"

	ginhandler "github.com/devpablocristo/golang/hex-arch/backend/internal/patients/infrastructure/handlers/gin"
)

func setupPatientsRoutes(r *gin.Engine, ph *ginhandler.GinHandler) {
	r.GET("/patient/:id", ph.GetPatient)
	r.POST("/patient", ph.CreatePatient)
}
