package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	event "github.com/devpablocristo/qh/internal/event-manager/event"
	ctypes "github.com/devpablocristo/qh/internal/platform/custom-types"
)

type Handler struct {
	useCase event.UseCasePort
}

type HandlerPort interface {
	CreateEvent(c *gin.Context)
	DeleteEvent(c *gin.Context)
	UpdateEvent(c *gin.Context)
	HardDeleteEvent(c *gin.Context)
	ReviveEvent(c *gin.Context)
	GetAllEvents(c *gin.Context)
	GetEvent(c *gin.Context)
	AddUserToEvent(c *gin.Context)
}

func NewHandler(uc event.UseCasePort) HandlerPort {
	return &Handler{
		useCase: uc,
	}
}

func (es *Handler) CreateEvent(c *gin.Context) {
	var dto *ctypes.EventDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
		return
	}
	_, err := es.useCase.CreateEvent(c, ctypes.EventDtoToDomain(dto))
	if err != nil {
		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, ctypes.NewAPIMessage("event successfully created"))
}

func (es *Handler) DeleteEvent(c *gin.Context) {
	eventID := c.Param("eventID")
	_, err := es.useCase.DeleteEvent(c, eventID)
	if err != nil {
		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, ctypes.NewAPIMessage("event successfully deleted"))
}

func (es *Handler) HardDeleteEvent(c *gin.Context) {
	eventID := c.Param("eventID")
	_, err := es.useCase.HardDeleteEvent(c, eventID)
	if err != nil {
		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, ctypes.NewAPIMessage("event successfully deleted"))
}

func (es *Handler) UpdateEvent(c *gin.Context) {
	var dto *ctypes.EventDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
		return
	}
	eventID := c.Param("eventID")
	_, err := es.useCase.UpdateEvent(c, ctypes.EventDtoToDomain(dto), eventID)
	if err != nil {
		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, ctypes.NewAPIMessage("event successfully updated"))
}

func (es *Handler) ReviveEvent(c *gin.Context) {
	eventID := c.Param("eventID")
	_, err := es.useCase.ReviveEvent(c, eventID)
	if err != nil {
		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, ctypes.NewAPIMessage("event successfully undeleted"))
}

func (es *Handler) GetEvent(c *gin.Context) {
	eventID := c.Param("eventID")
	event, err := es.useCase.GetEvent(c, eventID)
	if err != nil {
		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, ctypes.NewAPIMessage("event founded", event))
}

func (es *Handler) GetAllEvents(c *gin.Context) {
	events, err := es.useCase.GetAllEvents(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
		return
	}

	convEvents := event.EventToInterface(events)
	c.JSON(http.StatusOK, ctypes.NewAPIMessage("list of all events", convEvents))
}

func (es *Handler) AddUserToEvent(c *gin.Context) {
	eventID := c.Param("eventID")

	var dto *ctypes.UserDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
		return
	}

	_, err := es.useCase.AddUserToEvent(c, eventID, ctypes.UserDtoToDomain(dto))
	if err != nil {
		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, ctypes.NewAPIMessage("person added to event"))
}

// func (es *Handler) AddPersonsGroupToEvent(c *gin.Context) {
// 	events, err := es.useCase.GetAllEvents(c)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
// 		return
// 	}

// 	c.JSON(http.StatusOK, ctypes.NewAPIMessage("persons group to event"))

// }
