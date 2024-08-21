package ginpkgports

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Service interface {
	RunServer() error
	GetRouter() *gin.Engine
	WrapH(h http.Handler) gin.HandlerFunc
}

type Config interface {
	GetRouterPort() string
	SetRouterPort(string)
	Validate() error
}
