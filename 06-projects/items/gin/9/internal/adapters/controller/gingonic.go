// Para estar alineado con el naming de clean arch, se cambiar el nombre del directorio de controller a controller
package controller

import (
	"net/http"
	"strconv"

	// Se importa la librería Gin
	gin "github.com/gin-gonic/gin"

	"github.com/devpablocristo/golang/06-projects/items/gin/9/internal/entity"
	usecase "github.com/devpablocristo/golang/06-projects/items/gin/9/internal/usecase"
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
	var dto *itemDTO
	err := c.BindJSON(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item := dto2Item(dto)
	savedItems, err := h.usecase.SaveItem(item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, savedItems)
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

	c.JSON(http.StatusOK, items)
}

func (h *ItemController) GetItemsByID(c *gin.Context) {
	id := string2ID(c.Param(id))
	items, err := h.usecase.GetItemByID(id)
	if err != nil {
		if err == errNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, items)
}

func string2ID(s string) entity.ID {
	id, _ := strconv.Atoi(s)
	convID := entity.ID(id)
	return convID
}
