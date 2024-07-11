package handler

import (
	"context"

	"github.com/gin-gonic/gin"
)

type RouterService struct {
	eventHandler HandlerPort
}

type PortRouterService interface {
	SetupRouter(ctx context.Context) *gin.Engine
}

func NewRouterService(eh HandlerPort) PortRouterService {
	return &RouterService{
		eventHandler: eh,
	}
}

func (rs *RouterService) SetupRouter(ctx context.Context) *gin.Engine {
	r := gin.Default()

	r.POST("/events", rs.eventHandler.CreateEvent)
	r.DELETE("/events/:eventID", rs.eventHandler.DeleteEvent)
	r.DELETE("/events/hard/:eventID", rs.eventHandler.HardDeleteEvent)
	r.PATCH("/events/:eventID", rs.eventHandler.UpdateEvent)
	r.PATCH("/events/revive/:eventID", rs.eventHandler.ReviveEvent)
	r.GET("/events/:eventID", rs.eventHandler.GetEvent)
	r.GET("/events", rs.eventHandler.GetAllEvents)

	return r
}
