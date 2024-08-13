package portspkg

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GinClient interface {
	RunServer() error
	GetRouter() *gin.Engine
	WrapH(h http.Handler) gin.HandlerFunc
}

type GinConfig interface{}
