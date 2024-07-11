package gingonic

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GinClientPort interface {
	RunServer() error
	GetRouter() *gin.Engine
	WrapH(h http.Handler) gin.HandlerFunc
}
