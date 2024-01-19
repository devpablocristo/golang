package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"items/internal/domain"
	"items/internal/usecase"
)

type ItemHandler struct {
	usecase usecase.ItemUsecasePort
}

func NewHandler(u usecase.ItemUsecasePort) *ItemHandler {
	return &ItemHandler{
		usecase: u,
	}
}

func (h *ItemHandler) HelloWorld(c *gin.Context) {
	c.String(http.StatusOK, "¡Hello World!")
}

func (h *ItemHandler) SaveItem(c *gin.Context) {
	var item domain.Item
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	savedItem, err := h.usecase.SaveItem(item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, savedItem)
}

func (h *ItemHandler) GetAllItems(c *gin.Context) {
	items, err := h.usecase.GetAllItems()
	if err != nil {
		if err == domain.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, items)
}
