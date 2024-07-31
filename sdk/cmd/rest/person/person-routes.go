package handler

import (
	"context"

	"github.com/gin-gonic/gin"
)

type PersonRounter struct {
	eventHandler PersonHandlerPort
}

type PortRouterService interface {
	SetupRouter(ctx context.Context) *gin.Engine
}

func NewPersonRouter(eh PersonHandlerPort) PortRouterService {
	return &PersonRounter{
		eventHandler: eh,
	}
}

func (rs *PersonRounter) SetupRouter(ctx context.Context) *gin.Engine {
	r := gin.Default()

	r.POST("/events", rs.eventHandler.CreatePerson)
	// r.DELETE("/events/:eventID", rs.eventHandler.DeleteEvent)
	// r.DELETE("/events/hard/:eventID", rs.eventHandler.HardDeleteEvent)
	// r.PATCH("/events/:eventID", rs.eventHandler.UpdateEvent)
	// r.PATCH("/events/revive/:eventID", rs.eventHandler.ReviveEvent)
	// r.GET("/events/:eventID", rs.eventHandler.GetEvent)
	// r.GET("/events", rs.eventHandler.GetAllEvents)

	return r
}

func ChiRouter(handler *chihandler.ChiHandler) *chi.Mux {
	router := chi.NewRouter()
	//chiMux.Use("cors")
	//chiMux.Use(middleware.Logger)

	router.Route("/api/v1", func(r chi.Router) {
		r.Route("/person", func(r chi.Router) {
			r.Post("/create", handler.CreatePerson)
			r.Get("/list", handler.GetPersons)
			r.Get("/get/{personUUID}", handler.GetPerson)
			r.Put("/update/{personUUID}", handler.UpdatePerson)
			r.Delete("/delete", handler.DeletePerson)
		})
	})

	return router
}
