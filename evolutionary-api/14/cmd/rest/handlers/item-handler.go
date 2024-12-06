package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"api/internal/core"
	"api/internal/core/item"
	"api/pkg/config"
)

type handler struct {
	core core.ItemUsecasePort
}

func NewHandler(u core.ItemUsecasePort) *handler {
	return &handler{
		core: u,
	}
}

func (h *handler) SaveItem(c *gin.Context) {
	var item item.Item
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.core.SaveItem(item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "item saved successfully")
}

func (h *handler) ListItems(c *gin.Context) {
	items, err := h.core.ListItems()
	if err != nil {
		if err == config.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, items)
}
