package router

import (
	"github.com/gin-gonic/gin"

	ws "github.com/devpablocristo/golang/sdk/cmd/websocket/chat/wshandler"
)

func Routes(r *gin.Engine) {
	wshandler := ws.NewWsHandler()

	r.GET("/", wshandler.Home)
	r.GET("/ws", wshandler.WsEndpoint)
}
