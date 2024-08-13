package gtwports

import "github.com/gin-gonic/gin"

type GinHandler interface {
	CreateUser(c *gin.Context)
	GetUser(c *gin.Context)
	ListUsers(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}
