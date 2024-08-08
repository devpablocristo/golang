package wshandler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// upgradeConnection is the websocket upgrader from gorilla/websockets
var upgradeConnection = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// WsHandler estructura para manejar WebSocket con Gin
type WsHandler struct {
	upgrader websocket.Upgrader
}

// NewWsHandler crea una nueva instancia de WsHandler
func NewWsHandler() *WsHandler {
	return &WsHandler{
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				// Permitir todas las conexiones
				return true
			},
		},
	}
}

// Home maneja la solicitud de la página de inicio
func (h *WsHandler) Home(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to the HomePage!")
}

// Reader lee mensajes desde la conexión WebSocket
func (h *WsHandler) Reader(ws *websocket.Conn) {
	for {
		msgType, msg, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		log.Println(string(msg))

		if err := ws.WriteMessage(msgType, msg); err != nil {
			log.Println(err)
			return
		}
	}
}

// WsEndpoint maneja las solicitudes de WebSocket
func (h *WsHandler) WsEndpoint(c *gin.Context) {
	ws, err := h.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Fatal(err)
	}

	h.Reader(ws)
}
