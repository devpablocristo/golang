package chat

import (
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	wshandler := NewWsHandler()

	r.GET("/", wshandler.Home)
	r.GET("/ws", wshandler.WsEndpoint)
}
