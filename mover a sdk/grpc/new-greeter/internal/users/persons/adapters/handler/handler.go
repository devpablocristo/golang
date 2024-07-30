package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	ctypes "github.com/devpablocristo/qh/internal/platform/custom-types"
	port "github.com/devpablocristo/qh/internal/users/persons/ports"
)

type Handler struct {
	service port.Service
}

type PortHandler interface {
	CreatePerson(c *gin.Context)
	// DeletePerson(c *gin.Context)
	// UpdatePerson(c *gin.Context)
	// HardDeletePerson(c *gin.Context)
	// RevivePerson(c *gin.Context)
	// GetAllPersons(c *gin.Context)
	// GetPerson(c *gin.Context)
}

func NewHandler(es port.Service) PortHandler {
	return &Handler{
		service: es,
	}
}

func (es *Handler) CreatePerson(c *gin.Context) {
	// var dto *PersonDTO
	// if err := c.ShouldBindJSON(&dto); err != nil {
	// 	c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
	// 	return
	// }
	// _, err := es.service.CreatePerson(c, dtoToDomain(dto))
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
	// 	return
	// }

	c.JSON(http.StatusOK, ctypes.NewAPIMessage("person successfully created"))
}

// func (es *Handler) DeletePerson(c *gin.Context) {
// 	PersonID := c.Param("PersonID")
// 	_, err := es.service.DeletePerson(c, PersonID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
// 		return
// 	}

// 	c.JSON(http.StatusOK, ctypes.NewAPIMessage("Person successfully deleted"))
// }

// func (es *Handler) HardDeletePerson(c *gin.Context) {
// 	PersonID := c.Param("PersonID")
// 	_, err := es.service.HardDeletePerson(c, PersonID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
// 		return
// 	}

// 	c.JSON(http.StatusOK, ctypes.NewAPIMessage("Person successfully deleted"))
// }

// func (es *Handler) UpdatePerson(c *gin.Context) {
// 	var dto *PersonDTO
// 	if err := c.ShouldBindJSON(&dto); err != nil {
// 		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
// 		return
// 	}
// 	PersonID := c.Param("PersonID")
// 	_, err := es.service.UpdatePerson(c, dtoToDomain(dto), PersonID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
// 		return
// 	}

// 	c.JSON(http.StatusOK, ctypes.NewAPIMessage("Person successfully updated"))
// }

// func (es *Handler) RevivePerson(c *gin.Context) {
// 	PersonID := c.Param("PersonID")
// 	_, err := es.service.RevivePerson(c, PersonID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
// 		return
// 	}

// 	c.JSON(http.StatusOK, ctypes.NewAPIMessage("Person successfully undeleted"))
// }

// func (es *Handler) GetPerson(c *gin.Context) {
// 	PersonID := c.Param("PersonID")
// 	Person, err := es.service.GetPerson(c, PersonID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
// 		return
// 	}

// 	c.JSON(http.StatusOK, ctypes.NewAPIMessage("Person founded", Person))
// }

// func (es *Handler) GetAllPersons(c *gin.Context) {
// 	Persons, err := es.service.GetAllPersons(c)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
// 		return
// 	}

// 	convPersons := domain.PersonToInterface(Persons)
// 	c.JSON(http.StatusOK, ctypes.NewAPIMessage("list of all Persons", convPersons))
// }
