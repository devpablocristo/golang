package controller

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/items/internal/adapters/controller/presenter"
	"github.com/mercadolibre/items/internal/entity"
	"github.com/mercadolibre/items/internal/usecase"
)

type ItemController struct {
	itemUsecase usecase.ItemUsecase
}

func NewItemController(itemUsecase usecase.ItemUsecase) ItemController {
	return ItemController{
		itemUsecase: itemUsecase,
	}
}

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  "pong",
	})
}

func (ctrl ItemController) GetItems(c *gin.Context) {
	items, err := ctrl.itemUsecase.GetAllItems()
	c.JSON(http.StatusInternalServerError, presenter.ApiError{
		StatusCode: http.StatusInternalServerError,
		Message:    fmt.Sprintf("error getting items: %s", err.Error()),
	})

	c.JSON(http.StatusOK, presenter.ItemsResponse{
		Error: false,
		Data:  presenter.Items(items),
	})
}

type itemRequestDTO struct {
	Code   string `json:"code" binding:"required"`
	Author string `json:"author" binding:"required"`
	Title  string `json:"title" binding:"required"`
	Price  int    `json:"price" binding:"required"`
	Stock  int    `json:"stock" binding:"required"`
}

func (ctrl ItemController) AddItem(c *gin.Context) {
	var itemRequest itemRequestDTO

	if err := c.ShouldBindJSON(&itemRequest); err != nil {
		c.JSON(http.StatusBadRequest, presenter.ApiError{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("invalid json: %s", err.Error()),
		})
		return
	}

	item := entity.Item{
		Code:        itemRequest.Code,
		Description: itemRequest.Author,
		Title:       itemRequest.Title,
		Price:       itemRequest.Price,
		Stock:       itemRequest.Stock,
	}

	result, err := ctrl.itemUsecase.AddItem(item)
	if err != nil {
		var errorMsg string
		var httpStatus int

		existError := new(entity.ItemAlreadyExist)
		if ok := errors.As(err, existError); ok {
			httpStatus = http.StatusBadRequest
			errorMsg = existError.Error()
		} else {
			httpStatus = http.StatusInternalServerError
			errorMsg = err.Error()
		}

		c.JSON(httpStatus, presenter.ApiError{
			StatusCode: httpStatus,
			Message:    errorMsg,
		})
		return
	}

	c.JSON(http.StatusOK, presenter.ItemResponse{
		Error: false,
		Data:  presenter.Item(result),
	})
}

func (ctrl *ItemController) GetItem(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, presenter.ApiError{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("invalid param: %s", err.Error()),
		})
		return
	}

	item, err := ctrl.itemUsecase.GetItemByID(id)
	if err != nil {
		var errorMsg string
		var httpStatus int

		notFoundError := new(entity.ItemNotFound)
		if ok := errors.As(err, notFoundError); ok {
			httpStatus = http.StatusNotFound
			errorMsg = notFoundError.Error()
		} else {
			httpStatus = http.StatusInternalServerError
			errorMsg = err.Error()
		}

		c.JSON(httpStatus, presenter.ApiError{
			StatusCode: httpStatus,
			Message:    errorMsg,
		})
		return
	}

	c.JSON(http.StatusOK, presenter.ItemResponse{
		Error: false,
		Data:  presenter.Item(item),
	})
}
