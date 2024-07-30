package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	ctypes "github.com/devpablocristo/qh/internal/platform/custom-types"
	person "github.com/devpablocristo/qh/internal/user-manager/user/person"
)

type PersonHandler struct {
	service person.UseCasePort
}

type PersonHandlerPort interface {
	CreatePerson(c *gin.Context)
	// DeletePerson(c *gin.Context)
	// UpdatePerson(c *gin.Context)
	// HardDeletePerson(c *gin.Context)
	// RevivePerson(c *gin.Context)
	// GetAllPersons(c *gin.Context)
	// GetPerson(c *gin.Context)
}

func NewPersonHandler(es person.UseCasePort) PersonHandlerPort {
	return &PersonHandler{
		service: es,
	}
}

func (es *PersonHandler) CreatePerson(c *gin.Context) {
	// var dto *PersonDTO
	// if err := c.ShouldBindJSON(&dto); err != nil {
	// 	c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
	// 	return
	// }
	// _, err := es.UseCasePort.CreatePerson(c, dtoToDomain(dto))
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
	// 	return
	// }

	c.JSON(http.StatusOK, ctypes.NewAPIMessage("person successfully created"))
}

// func (es *PersonHandler) DeletePerson(c *gin.Context) {
// 	PersonID := c.Param("PersonID")
// 	_, err := es.UseCasePort.DeletePerson(c, PersonID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
// 		return
// 	}

// 	c.JSON(http.StatusOK, ctypes.NewAPIMessage("Person successfully deleted"))
// }

// func (es *PersonHandler) HardDeletePerson(c *gin.Context) {
// 	PersonID := c.Param("PersonID")
// 	_, err := es.UseCasePort.HardDeletePerson(c, PersonID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
// 		return
// 	}

// 	c.JSON(http.StatusOK, ctypes.NewAPIMessage("Person successfully deleted"))
// }

// func (es *PersonHandler) UpdatePerson(c *gin.Context) {
// 	var dto *PersonDTO
// 	if err := c.ShouldBindJSON(&dto); err != nil {
// 		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
// 		return
// 	}
// 	PersonID := c.Param("PersonID")
// 	_, err := es.UseCasePort.UpdatePerson(c, dtoToDomain(dto), PersonID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
// 		return
// 	}

// 	c.JSON(http.StatusOK, ctypes.NewAPIMessage("Person successfully updated"))
// }

// func (es *PersonHandler) RevivePerson(c *gin.Context) {
// 	PersonID := c.Param("PersonID")
// 	_, err := es.UseCasePort.RevivePerson(c, PersonID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
// 		return
// 	}

// 	c.JSON(http.StatusOK, ctypes.NewAPIMessage("Person successfully undeleted"))
// }

// func (es *PersonHandler) GetPerson(c *gin.Context) {
// 	PersonID := c.Param("PersonID")
// 	Person, err := es.UseCasePort.GetPerson(c, PersonID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
// 		return
// 	}

// 	c.JSON(http.StatusOK, ctypes.NewAPIMessage("Person founded", Person))
// }

// func (es *PersonHandler) GetAllPersons(c *gin.Context) {
// 	Persons, err := es.UseCasePort.GetAllPersons(c)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, ctypes.NewAPIError(http.StatusBadRequest, err.Error()))
// 		return
// 	}

// 	convPersons := domain.PersonToInterface(Persons)
// 	c.JSON(http.StatusOK, ctypes.NewAPIMessage("list of all Persons", convPersons))
// }
