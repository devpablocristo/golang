// Para estar alineado con el naming de clean arch, se cambiar el nombre del directorio de controller a controller
package controller

import (
	"net/http"
	"strconv"

	// Se importa la librería Gin
	gin "github.com/gin-gonic/gin"

	presenter "items/internal/adapters/controller/presenter"
	entity "items/internal/entity"
	usecase "items/internal/usecase"
)

// ATENCION aqui se ultiliza la interface del usercase, no el tipo del usercase
// interactor ques la estructura?????
type ItemController struct {
	usecase usecase.ItemUsecaseInterface
}

// Constructor del tipo ItemController, en los parametros de entrada se inyecta el un usecase
// como el campo usecase es de tipo interface, tiene sentido poner como paramtro de entrada tambien la misma interface
func NewController(u usecase.ItemUsecaseInterface) *ItemController {
	return &ItemController{
		usecase: u, // Aquí se carga el usecase inyectado dentro del ItemController
	}
}

func (h *ItemController) SaveItem(c *gin.Context) {
	var dto itemDTO         // Declarar una variable de tipo itemDTO
	err := c.BindJSON(&dto) // Pasar la dirección de dto a BindJSON
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error1": err.Error()})
		return
	}

	item := dto.dto2Item() // Pasar la dirección de dto a dto2Item
	savedItem, err := h.usecase.SaveItem(item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error2": err.Error()})
		return
	}

	c.JSON(http.StatusOK, presenter.Item(savedItem))
}

// devuelve un array de items
func (h *ItemController) GetAllItems(c *gin.Context) {
	items, err := h.usecase.GetAllItems()
	if err != nil {
		if err == errNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, presenter.Items(items))
}

func (h *ItemController) GetItemsByID(c *gin.Context) {
	id := string2ID(c.Param("id"))
	item, err := h.usecase.GetItemByID(id)
	if err != nil {
		if err == errNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, presenter.Item(item))
}

func string2ID(s string) entity.ID {
	id, _ := strconv.Atoi(s)
	convID := entity.ID(id)
	return convID
}
