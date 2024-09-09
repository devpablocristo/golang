package event

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"

	shared "github.com/devpablocristo/golang/sdk/cmd/gateways/shared"
	coreeventports "github.com/devpablocristo/golang/sdk/internal/core/event/ports"
	sdkgin "github.com/devpablocristo/golang/sdk/pkg/rest/gin/ports"
)

type Handler struct {
	useCases  coreeventports.UseCases
	ginServer sdkgin.Server
}

// NewHandler crea un nuevo handler para los eventos.
func NewHandler(uc coreeventports.UseCases, gs sdkgin.Server) *Handler {
	return &Handler{
		useCases:  uc,
		ginServer: gs,
	}
}

func (h *Handler) Start(apiVersion string, secret string) error {
	h.Routes(apiVersion, secret)
	return h.ginServer.RunServer()
}

func (h *Handler) Routes(apiVersion string, secret string) {
	//r := h.ginServer.GetRouter()

	// r.POST("/events", r.CreateEvent)
	// r.DELETE("/events/:eventID", r.DeleteEvent)
	// r.DELETE("/events/hard/:eventID", r.HardDeleteEvent)
	// r.PATCH("/events/:eventID", r.UpdateEvent)
	// r.PATCH("/events/revive/:eventID", r.ReviveEvent)
	// r.GET("/events/:eventID", r.GetEvent)
	// r.GET("/events", r.GetAllEvents)
}

// FakeCreateEvent maneja la creación ficticia de un evento para pruebas.
func (h *Handler) FakeCreateEvent(c *gin.Context) {
	dir := "data"
	filename := "event.json"

	data, err := shared.LoadTestData(dir, filename)
	if err != nil {
		shared.WriteErrorResponse(c.Writer, shared.ApiErrors["InternalServer"], c.Request.Method)
		return
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(data))

	req, err := decodeCreateEventRequest(c)
	if err != nil {
		shared.WriteErrorResponse(c.Writer, shared.ApiErrors["InvalidJSON"], c.Request.Method)
		return
	}

	if err = h.useCases.CreateEvent(c.Request.Context(), req.ToDomain()); err != nil {
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
