package ginhandler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	patient "github.com/devpablocristo/golang/hex-arch/backend/internal/patients/domain"
	ports "github.com/devpablocristo/golang/hex-arch/backend/internal/patients/service/ports"
)

type GinHandler struct {
	patientService  ports.Service
	patientScrapper ports.Scrapper
}

// func NewGinHandler(ps ports.PatientService, gr *gin.Engine) *GinHandler {
func NewGinHandler(ps ports.Service, pe ports.Scrapper) *GinHandler {
	return &GinHandler{
		patientService:  ps,
		patientScrapper: pe,
	}
}

// func (gh *GinHandler) Run(port string) {
// 	if port == "default" {
// 		port = ":8000"
// 	}

// 	log.Printf("Gin Server listening on %s\n", port)
// 	log.Fatal(http.ListenAndServe(port, gh.ginRouter))
// }

// func (gh *GinHandler) SetupRoutes() {
// 	gh.ginRouter.GET("/patient/:id", func(c *gin.Context) {
// 		gh.GetPatient(c)
// 	})
// 	gh.ginRouter.POST("/patient", func(c *gin.Context) {
// 		gh.CreatePatient(c)
// 	})
// }

func (gh *GinHandler) GetPatient(c *gin.Context) {
	c.Set("content-type", "application/json")

	// path param patients/1234
	fmt.Printf("%T", c.Param("id"))
	patient, err := gh.patientService.GetPatient(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": patient,
	})
}

func (gh *GinHandler) CreatePatient(c *gin.Context) {
	c.Set("content-type", "application/json")

	newPatient := patient.Patient{}
	//<c.ShouldBindJSON(&newPatient)
	c.BindJSON(&newPatient)

	patient, err := gh.patientService.CreatePatient(newPatient)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": patient,
	})
}
