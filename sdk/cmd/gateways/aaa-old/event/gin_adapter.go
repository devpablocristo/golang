package event

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/devpablocristo/golang/sdk/cmd/gateways/shared"
	"github.com/devpablocristo/golang/sdk/cmd/gateways/ttools"
	"github.com/devpablocristo/golang/sdk/internal/core"
)

type Handler struct {
	uc core.EventUseCases
}

// NewHandler crea un nuevo handler para los eventos.
func NewHandler(uc core.EventUseCases) *Handler {
	return &Handler{
		uc: uc,
	}
}

// FakeCreateEvent maneja la creación ficticia de un evento para pruebas.
func (h *Handler) FakeCreateEvent(c *gin.Context) {
	data, err := ttools.LoadTestData("event.json")
	if err != nil {
		shared.WriteErrorResponse(c.Writer, shared.ApiErrors["InternalServer"], c.Request.Method)
		return
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(data))

	dto, err := decodeCreateEventRequest(c)
	if err != nil {
		shared.WriteErrorResponse(c.Writer, shared.ApiErrors["InvalidJSON"], c.Request.Method)
		return
	}

	if err = h.uc.CreateEvent(c.Request.Context(), dto.ToDomain()); err != nil {
		shared.WriteErrorResponse(c.Writer, shared.ApiErrors["InternalServer"], c.Request.Method)
		return
	}

	response := shared.NewApiResponse(true, http.StatusOK, "Event created successfully", nil)
	shared.WriteJSONResponse(c.Writer, http.StatusOK, response)
}

// Health maneja la verificación de estado de la API.
func (h *Handler) Health(c *gin.Context) {
	response := shared.NewApiResponse(true, http.StatusOK, "Service is UP", nil)
	shared.WriteJSONResponse(c.Writer, http.StatusOK, response)
}

// func (es *Handler) CreateEvent(c *gin.Context) {
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

// func (es *Handler) DeleteEvent(c *gin.Context) {
// 	eventID := c.Param("eventID")
// 	_, err := es.useCase.DeleteEvent(c, eventID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
// 		return
// 	}

// 	c.JSON(http.StatusOK, ctypes.NewAPIMessage("event successfully deleted"))
// }

// func (es *Handler) HardDeleteEvent(c *gin.Context) {
// 	eventID := c.Param("eventID")
// 	_, err := es.useCase.HardDeleteEvent(c, eventID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
// 		return
// 	}

// 	c.JSON(http.StatusOK, ctypes.NewAPIMessage("event successfully deleted"))
// }

// func (es *Handler) UpdateEvent(c *gin.Context) {
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

// func (es *Handler) ReviveEvent(c *gin.Context) {
// 	eventID := c.Param("eventID")
// 	_, err := es.useCase.ReviveEvent(c, eventID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
// 		return
// 	}

// 	c.JSON(http.StatusOK, ctypes.NewAPIMessage("event successfully undeleted"))
// }

// func (es *Handler) GetEvent(c *gin.Context) {
// 	eventID := c.Param("eventID")
// 	event, err := es.useCase.GetEvent(c, eventID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
// 		return
// 	}

// 	c.JSON(http.StatusOK, ctypes.NewAPIMessage("event founded", event))
// }

// func (es *Handler) GetAllEvents(c *gin.Context) {
// 	events, err := es.useCase.GetAllEvents(c)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
// 		return
// 	}

// 	convEvents := event.EventToInterface(events)
// 	c.JSON(http.StatusOK, ctypes.NewAPIMessage("list of all events", convEvents))
// }

// func (es *Handler) AddUserToEvent(c *gin.Context) {
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

// // func (es *Handler) AddPersonsGroupToEvent(c *gin.Context) {
// // 	events, err := es.useCase.GetAllEvents(c)
// // 	if err != nil {
// // 		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
// // 		return
// // 	}

// // 	c.JSON(http.StatusOK, ctypes.NewAPIMessage("persons group to event"))

// // }

// type RestHandler struct {
// 	core uc.UseCasesPort
// }

// func NewRestHandler(core uc.UseCasesPort) *RestHandler {
// 	return &RestHandler{
// 		core: core,
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

// 	if err = h.uc.CreateEvent(c.Request.Context(), dto.ToDomain()); err != nil {
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