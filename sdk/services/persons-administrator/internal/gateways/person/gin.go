package person

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/devpablocristo/golang/sdk/services/person/gateways/person/dto"
	"github.com/devpablocristo/golang/sdk/services/person/internal/person/ports"
	"github.com/devpablocristo/golang/sdk/services/shared"
)

type GinHandler struct {
	service ports.UseCases
}

// NewGinHandler crea una nueva instancia de GinHandler.
func NewGinHandler(service ports.UseCases) *GinHandler {
	return &GinHandler{
		service: service,
	}
}

func GinRoutes(r *gin.Engine) {
	// r := gin.Default()

	// r.POST("/events", rs.eventHandler.CreatePerson)
	// r.DELETE("/events/:eventID", rs.eventHandler.DeleteEvent)
	// r.DELETE("/events/hard/:eventID", rs.eventHandler.HardDeleteEvent)
	// r.PATCH("/events/:eventID", rs.eventHandler.UpdateEvent)
	// r.PATCH("/events/revive/:eventID", rs.eventHandler.ReviveEvent)
	// r.GET("/events/:eventID", rs.eventHandler.GetEvent)
	// r.GET("/events", rs.eventHandler.GetAllEvents)

	// return r
}

func (h *GinHandler) CreatePerson(c *gin.Context) {
	var dto dto.PersonRequest
	if err := c.ShouldBindJSON(&dto); err != nil {
		shared.WriteErrorResponse(c.Writer, shared.ApiErrors["BadRequest"], "GinHandler.CreatePerson")
		return
	}

	if err := h.service.CreatePerson(c.Request.Context(), dto.ToDomain()); err != nil {
		shared.WriteErrorResponse(c.Writer, shared.ApiErrors["InternalServer"], "GinHandler.CreatePerson")
		return
	}

	response := shared.NewApiResponse(true, http.StatusCreated, "Person created successfully", dto.ToDomain())
	shared.WriteJSONResponse(c.Writer, http.StatusCreated, response)
}

// func (es *GinHandler) DeletePerson(c *gin.Context) {
// 	PersonID := c.Param("PersonID")
// 	_, err := es.PersonUseCasesPort.DeletePerson(c, PersonID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
// 		return
// 	}

// 	c.JSON(http.StatusOK, ctypes.NewAPIMessage("Person successfully deleted"))
// }

// func (es *GinHandler) HardDeletePerson(c *gin.Context) {
// 	PersonID := c.Param("PersonID")
// 	_, err := es.PersonUseCasesPort.HardDeletePerson(c, PersonID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
// 		return
// 	}

// 	c.JSON(http.StatusOK, ctypes.NewAPIMessage("Person successfully deleted"))
// }

// func (es *GinHandler) UpdatePerson(c *gin.Context) {
// 	var dto *dto.PersonRequest
// 	if err := c.ShouldBindJSON(&dto); err != nil {
// 		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
// 		return
// 	}
// 	PersonID := c.Param("PersonID")
// 	_, err := es.PersonUseCasesPort.UpdatePerson(c, dtoToDomain(dto), PersonID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
// 		return
// 	}

// 	c.JSON(http.StatusOK, ctypes.NewAPIMessage("Person successfully updated"))
// }

// func (es *GinHandler) RevivePerson(c *gin.Context) {
// 	PersonID := c.Param("PersonID")
// 	_, err := es.PersonUseCasesPort.RevivePerson(c, PersonID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
// 		return
// 	}

// 	c.JSON(http.StatusOK, ctypes.NewAPIMessage("Person successfully undeleted"))
// }

// func (es *GinHandler) GetPerson(c *gin.Context) {
// 	PersonID := c.Param("PersonID")
// 	Person, err := es.PersonUseCasesPort.GetPerson(c, PersonID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
// 		return
// 	}

// 	c.JSON(http.StatusOK, ctypes.NewAPIMessage("Person founded", Person))
// }
