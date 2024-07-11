package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"api/internal/domain"
	usecase "api/internal/usecases"
)

type handler struct {
	usecase usecase.ItemUsecasePort
}

func NewHandler(u usecase.ItemUsecasePort) *handler {
	return &handler{
		usecase: u,
	}
}

func (h *handler) SaveItem(c *gin.Context) {
	var it domain.Item
	err := c.BindJSON(&it)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.usecase.SaveItem(it); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "item saved successfully")
}

func (h *handler) ListItems(c *gin.Context) {
	its, err := h.usecase.ListItems()
	if err != nil {
		if err == domain.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, its)
}
