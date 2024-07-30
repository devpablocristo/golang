package handler

import (
	"fmt"
	"net/http"
	"os/user"
	"strconv"

	"github.com/devpablocristo/golang/06-projects/qh/person/application/port"

	"github.com/gin-gonic/gin"
)

type GinHandler struct {
	userService port.Service
}

func NewGinHandler(us port.Service) *GinHandler {
	return &GinHandler{
		userService: us,
	}
}

func CreateUser(c *gin.Context) {
	var user domain.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		restErr := errors.BadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		fmt.Println(restErr)
		return
	}

	result, rErr := service.CreateUser(user)
	if err != nil {
		c.JSON(rErr.Status, rErr)
		return
	}

	c.JSON(http.StatusCreated, result)

	// var input = inputRequest{}
	// if err := ctx.ShouldBindJSON(&input); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, err.Error())
	// 	return
	// }

	// redirect, err := useCase.UrlToCode(input.Url)
	// if err != nil {
	// 	if err == entity.ErrRedirectInvalid {
	// 		ctx.JSON(http.StatusBadRequest, err.Error())
	// 		return
	// 	}
	// 	ctx.JSON(http.StatusInternalServerError, err.Error())
	// 	return
	// }
	// ctx.JSON(http.StatusOK, redirect)

}

func (h *GinHandler) GetUsers(c *gin.Context) {
	result, rErr := service.GetUsers()
	if rErr != nil {
		c.JSON(rErr.Status, rErr)
		return
	}

	c.JSON(http.StatusCreated, result)

}

func (h *GinHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	result, rErr := service.GetUser(id)
	if rErr != nil {
		c.JSON(rErr.Status, rErr)
		return
	}

	c.JSON(http.StatusCreated, result)

	// code := ctx.Param("code")
	// redirect, err := useCase.CodeToUrl(code)
	// if err != nil {
	// 	if err == entity.ErrRedirectNotFound {
	// 		ctx.JSON(http.StatusNotFound, err.Error())
	// 		return
	// 	}
	// 	ctx.JSON(http.StatusInternalServerError, err.Error())
	// 	return
	// }
	// ctx.Redirect(http.StatusPermanentRedirect, redirect.URL)
}

func (h *GinHandler) updateUser(c *gin.Context) {
	var u user.User

	err := c.ShouldBindJSON(&u)
	if err != nil {
		restErr := errors.BadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		fmt.Println(restErr)
		return
	}

	/*
		ejemplo para actualizar

		id: 609a4ad964b4678593e14d6a
		{
			"username":"LedZeppelin",
			"password":"rock",
			"email":"musica@rock.com"
		}

	*/

	uId := c.Param("id")
	result, rErr := service.UpdateUser(u, uId)
	if rErr != nil {
		c.JSON(rErr.Status, rErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (h *GinHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	del, rErr := service.DeleteUser(id)
	if rErr != nil {
		c.JSON(rErr.Status, rErr)
		return
	}
	nDel := int(*del)
	r := "Deleted " + strconv.Itoa(nDel) + " document/s."

	c.JSON(http.StatusCreated, r)
}
