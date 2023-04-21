package controller

import (
	"Items/internal/application"
	entity "Items/internal/domain"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	ucService application.UseCase
}

func NewController(useService application.UseCase) Controller {

	controllerService := Controller{
		ucService: useService,
	}
	return controllerService
}

func (c *Controller) AddItem(ctx *gin.Context) {
	item := entity.Item{}
	c.ucService.AddItem(item)
}
