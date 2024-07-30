package users

import (
	"fmt"
	"net/http"
	"strconv"

	users "github.com/devpablocristo/bookstore_users_api.dao/domain/users"
	service "github.com/devpablocristo/bookstore_users_api.dao/services"
	errors "github.com/devpablocristo/bookstore_users_api.dao/utils/errors"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User
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
}

/*
func GetUsers(c *gin.Context) {
	result, rErr := service.GetUsers()
	if rErr != nil {
		c.JSON(rErr.Status, rErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}
*/

func GetUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		rErr := errors.BadRequestError("invalida user id")
		c.JSON(rErr.Status, rErr)
	}

	result, rErr := service.GetUser(id)
	if rErr != nil {
		c.JSON(rErr.Status, rErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

/*
func UpdateUser(c *gin.Context) {
	var u users.User

	err := c.ShouldBindJSON(&u)
	if err != nil {
		restErr := errors.BadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		fmt.Println(restErr)
		return
	}
*/
/*
	ejemplo para actualizar

	id: 609a4ad964b4678593e14d6a
	{
		"username":"LedZeppelin",
		"password":"rock",
		"email":"musica@rock.com"
	}

*/
/*
	uId := c.Param("id")
	result, rErr := service.UpdateUser(u, uId)
	if rErr != nil {
		c.JSON(rErr.Status, rErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func DeleteUser(c *gin.Context) {
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
*/
