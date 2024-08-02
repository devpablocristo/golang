package person

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/devpablocristo/golang/sdk/cmd/rest/shared"
	"github.com/devpablocristo/golang/sdk/internal/core"
)

type GinHandler struct {
	service core.PersonUseCasePort
}

type GinHandlerPort interface {
	CreatePerson(c *gin.Context)
	// DeletePerson(c *gin.Context)
	// UpdatePerson(c *gin.Context)
	// HardDeletePerson(c *gin.Context)
	// RevivePerson(c *gin.Context)
	// GetAllPersons(c *gin.Context)
	// GetPerson(c *gin.Context)
}

// NewGinHandler crea una nueva instancia de GinHandler.
func NewGinHandler(service core.PersonUseCasePort) GinHandlerPort {
	return &GinHandler{
		service: service,
	}
}

// CreatePerson maneja la creación de una nueva persona utilizando Gin.
func (h *GinHandler) CreatePerson(c *gin.Context) {
	var dto PersonDTO

	// Validar la entrada JSON en la estructura DTO.
	if err := c.ShouldBindJSON(&dto); err != nil {
		shared.WriteErrorResponse(c.Writer, shared.ApiErrors["BadRequest"], "GinHandler.CreatePerson")
		return
	}

	// Convertir DTO a modelo de dominio.
	newPerson := dto.ToDomain()

	// Intentar crear la persona en la base de datos.
	ctx := c.Request.Context()
	if err := h.service.CreatePerson(ctx, newPerson); err != nil {
		shared.WriteErrorResponse(c.Writer, shared.ApiErrors["InternalServer"], "GinHandler.CreatePerson")
		return
	}

	// Responder con un mensaje de éxito.
	response := shared.NewApiResponse(true, http.StatusCreated, "Person created successfully", newPerson)
	shared.WriteJSONResponse(c.Writer, http.StatusCreated, response)
}

// func (es *GinHandler) DeletePerson(c *gin.Context) {
// 	PersonID := c.Param("PersonID")
// 	_, err := es.PersonUseCasePort.DeletePerson(c, PersonID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
// 		return
// 	}

// 	c.JSON(http.StatusOK, ctypes.NewAPIMessage("Person successfully deleted"))
// }

// func (es *GinHandler) HardDeletePerson(c *gin.Context) {
// 	PersonID := c.Param("PersonID")
// 	_, err := es.PersonUseCasePort.HardDeletePerson(c, PersonID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
// 		return
// 	}

// 	c.JSON(http.StatusOK, ctypes.NewAPIMessage("Person successfully deleted"))
// }

// func (es *GinHandler) UpdatePerson(c *gin.Context) {
// 	var dto *PersonDTO
// 	if err := c.ShouldBindJSON(&dto); err != nil {
// 		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
// 		return
// 	}
// 	PersonID := c.Param("PersonID")
// 	_, err := es.PersonUseCasePort.UpdatePerson(c, dtoToDomain(dto), PersonID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
// 		return
// 	}

// 	c.JSON(http.StatusOK, ctypes.NewAPIMessage("Person successfully updated"))
// }

// func (es *GinHandler) RevivePerson(c *gin.Context) {
// 	PersonID := c.Param("PersonID")
// 	_, err := es.PersonUseCasePort.RevivePerson(c, PersonID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
// 		return
// 	}

// 	c.JSON(http.StatusOK, ctypes.NewAPIMessage("Person successfully undeleted"))
// }

// func (es *GinHandler) GetPerson(c *gin.Context) {
// 	PersonID := c.Param("PersonID")
// 	Person, err := es.PersonUseCasePort.GetPerson(c, PersonID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
// 		return
// 	}

// 	c.JSON(http.StatusOK, ctypes.NewAPIMessage("Person founded", Person))
// }

// func (es *GinHandler) GetAllPersons(c *gin.Context) {
// 	Persons, err := es.PersonUseCasePort.GetAllPersons(c)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
// 		return
// 	}

// 	convPersons := domain.PersonToInterface(Persons)
// 	c.JSON(http.StatusOK, ctypes.NewAPIMessage("list of all Persons", convPersons))
// }
