package ginAdapter

import (
	"log"
	"net/http"

	"github.com/devpablocristo/golang/hex-arch/backend/internal/persons/domain"
	"github.com/gin-gonic/gin"
)

type GinHandler struct {
	Service   ports.PersonService
	GinRouter *gin.Engine
}

func NewGinHandler(s ports.PersonService) *GinHandler {
	//gin.SetMode(gin.ReleaseMode)
	//r := gin.Default()
	return &GinHandler{
		Service:   s,
		GinRouter: r,
	}
}

func (h *GinHandler) Run(port string) {
	if port == "default" {
		port = ":8000"
	}

	log.Printf("Gin Server listening on %s\n", port)
	log.Fatal(http.ListenAndServe(port, h.GinRouter))
}

func (h *GinHandler) SetupRoutes() {
	h.GinRouter.GET("/person", func(ctx *gin.Context) {
		h.List(ctx)
	})
	h.GinRouter.POST("/person", func(ctx *gin.Context) {
		h.Register(ctx)
	})
}

func (h *GinHandler) List(ctx *gin.Context) {
	ctx.Set("content-type", "application/json")

	personsList := h.Service.List()
	ctx.JSON(http.StatusOK, personsList)
}

func (h *GinHandler) Register(ctx *gin.Context) {
	ctx.Set("content-type", "application/json")

	newPerson := domain.Person{}
	ctx.ShouldBindJSON(&newPerson)

	ctx.JSON(http.StatusOK, gin.H{
		"data": newPerson,
	})
}
