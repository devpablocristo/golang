package handler

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"

	tst "github.com/devpablocristo/golang/sdk/cmd/rest/tests"
	ucs "github.com/devpablocristo/golang/sdk/internal/core"
)

type ItemHandler struct {
	ucs ucs.EventUseCasePort
}

func NewItemHandler(ucs ucs.EventUseCasePort) *ItemHandler {
	return &ItemHandler{
		ucs: ucs,
	}
}

func (h *ItemHandler) FakeCreateEvent(c *gin.Context) {
	data, err := tst.LoadTestData()
	if err != nil {
		errorHandler(c, err)
		return
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(data))

	dto, err := decodeCreateEventRequest(c)
	if err != nil {
		errorHandler(c, err)
		return
	}

	if err = h.ucs.CreateEvent(c.Request.Context(), dto.ToDomain()); err != nil {
		errorHandler(c, err)
		return
	}

	response := CreateEventResponse{
		Message: "Event created successfully",
	}
	encodeResponse(c, response)
}

func (h *ItemHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "UP",
	})
}

// type EventHandler struct {
// 	useCase event.UseCasePort
// }

// type EventHandlerPort interface {
// 	CreateEvent(c *gin.Context)
// 	DeleteEvent(c *gin.Context)
// 	UpdateEvent(c *gin.Context)
// 	HardDeleteEvent(c *gin.Context)
// 	ReviveEvent(c *gin.Context)
// 	GetAllEvents(c *gin.Context)
// 	GetEvent(c *gin.Context)
// 	AddUserToEvent(c *gin.Context)
// }

// func NewEventHandler(uc event.UseCasePort) EventHandlerPort {
// 	return &EventHandler{
// 		useCase: uc,
// 	}
// }

// func (es *EventHandler) CreateEvent(c *gin.Context) {
// 	var dto *ctypes.EventDTO
// 	if err := c.ShouldBindJSON(&dto); err != nil {
// 		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
// 		return
// 	}
// 	_, err := es.useCase.CreateEvent(c, ctypes.EventDtoToDomain(dto))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
// 		return
// 	}

// 	c.JSON(http.StatusOK, ctypes.NewAPIMessage("event successfully created"))
// }

// func (es *EventHandler) DeleteEvent(c *gin.Context) {
// 	eventID := c.Param("eventID")
// 	_, err := es.useCase.DeleteEvent(c, eventID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
// 		return
// 	}

// 	c.JSON(http.StatusOK, ctypes.NewAPIMessage("event successfully deleted"))
// }

// func (es *EventHandler) HardDeleteEvent(c *gin.Context) {
// 	eventID := c.Param("eventID")
// 	_, err := es.useCase.HardDeleteEvent(c, eventID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
// 		return
// 	}

// 	c.JSON(http.StatusOK, ctypes.NewAPIMessage("event successfully deleted"))
// }

// func (es *EventHandler) UpdateEvent(c *gin.Context) {
// 	var dto *ctypes.EventDTO
// 	if err := c.ShouldBindJSON(&dto); err != nil {
// 		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
// 		return
// 	}
// 	eventID := c.Param("eventID")
// 	_, err := es.useCase.UpdateEvent(c, ctypes.EventDtoToDomain(dto), eventID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
// 		return
// 	}

// 	c.JSON(http.StatusOK, ctypes.NewAPIMessage("event successfully updated"))
// }

// func (es *EventHandler) ReviveEvent(c *gin.Context) {
// 	eventID := c.Param("eventID")
// 	_, err := es.useCase.ReviveEvent(c, eventID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
// 		return
// 	}

// 	c.JSON(http.StatusOK, ctypes.NewAPIMessage("event successfully undeleted"))
// }

// func (es *EventHandler) GetEvent(c *gin.Context) {
// 	eventID := c.Param("eventID")
// 	event, err := es.useCase.GetEvent(c, eventID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
// 		return
// 	}

// 	c.JSON(http.StatusOK, ctypes.NewAPIMessage("event founded", event))
// }

// func (es *EventHandler) GetAllEvents(c *gin.Context) {
// 	events, err := es.useCase.GetAllEvents(c)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
// 		return
// 	}

// 	convEvents := event.EventToInterface(events)
// 	c.JSON(http.StatusOK, ctypes.NewAPIMessage("list of all events", convEvents))
// }

// func (es *EventHandler) AddUserToEvent(c *gin.Context) {
// 	eventID := c.Param("eventID")

// 	var dto *ctypes.UserDTO
// 	if err := c.ShouldBindJSON(&dto); err != nil {
// 		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
// 		return
// 	}

// 	_, err := es.useCase.AddUserToEvent(c, eventID, ctypes.UserDtoToDomain(dto))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
// 		return
// 	}

// 	c.JSON(http.StatusOK, ctypes.NewAPIMessage("person added to event"))
// }

// // func (es *EventHandler) AddPersonsGroupToEvent(c *gin.Context) {
// // 	events, err := es.useCase.GetAllEvents(c)
// // 	if err != nil {
// // 		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
// // 		return
// // 	}

// // 	c.JSON(http.StatusOK, ctypes.NewAPIMessage("persons group to event"))

// // }

// type RestHandler struct {
// 	ucs ucs.UseCasePort
// }

// func NewRestHandler(ucs ucs.UseCasePort) *RestHandler {
// 	return &RestHandler{
// 		ucs: ucs,
// 	}
// }

// func (h *RestHandler) FakeCreateEvent(c *gin.Context) {
// 	data, err := tst.LoadTestData()
// 	if err != nil {
// 		errorHandler(c, err)
// 		return
// 	}
// 	c.Request.Body = io.NopCloser(bytes.NewBuffer(data))

// 	dto, err := decodeCreateEventRequest(c)
// 	if err != nil {
// 		errorHandler(c, err)
// 		return
// 	}

// 	if err = h.ucs.CreateEvent(c.Request.Context(), dto.ToDomain()); err != nil {
// 		errorHandler(c, err)
// 		return
// 	}

// 	response := CreateEventResponse{
// 		Message: "Event created successfully",
// 	}
// 	encodeResponse(c, response)
// }

// func (h *RestHandler) Health(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"status": "UP",
// 	})
// }
