package handler

import (
	"net/http"
	"strconv"

	gin "github.com/gin-gonic/gin"

	presenter "items/internal/adapters/handler/presenter"
	domain "items/internal/domain"
	ctypes "items/internal/platform/custom-types"
	usecase "items/internal/usecase"
)

type ItemHandler struct {
	usecase usecase.ItemUsecasePort
}

// NewHandler creates a new instance of ItemHandler with the given usecase.
func NewHandler(u usecase.ItemUsecasePort) *ItemHandler {
	return &ItemHandler{
		usecase: u,
	}
}

// SaveItem handles the request to save an item.
func (h *ItemHandler) SaveItem(c *gin.Context) {
	var dto itemDTO
	err := c.BindJSON(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item := dtoToItem(&dto)
	savedItem, err := h.usecase.SaveItem(item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, presenter.Item(savedItem))
}

// GetAllItems returns all items.
func (h *ItemHandler) GetAllItems(c *gin.Context) {
	items, err := h.usecase.GetAllItems()
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, presenter.Items(items))
}

// GetItem returns an item by its ID.
func (h *ItemHandler) GetItem(c *gin.Context) {
	id, err := stringToID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	item, err := h.usecase.GetItem(id)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, presenter.Item(item))
}

// stringToID converts a string to domain.ID, handling any conversion errors.
func stringToID(s string) (domain.ID, error) {
	id, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return domain.ID(id), nil
}

// handleError centralizes error handling for API responses.
func handleError(c *gin.Context, err error) {
	if err.Error() == ctypes.ErrItemNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
