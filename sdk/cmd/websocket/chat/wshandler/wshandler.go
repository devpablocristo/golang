package ws

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	chat "github.com/devpablocristo/interviews/b6/chat/domain"
	renderjet "github.com/devpablocristo/interviews/b6/chat/infrastructure/renderjet"
	usecases "github.com/devpablocristo/interviews/b6/chat/usecases"
)

// upgradeConnection is the websocket upgrader from gorilla/websockets
var upgradeConnection = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// WsEndpoint upgrades connection to websocket
func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgradeConnection.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client connected to endpoint")

	var response chat.WsJsonResponse
	response.Message = `<em><small>Connected to server</small></em>`

	conn := chat.WebSocketConnection{Conn: ws}
	usecases.Clients[conn] = ""

	err = ws.WriteJSON(response)
	if err != nil {
		log.Println(err)
	}

	go usecases.ListenForWs(&conn)
}

// Home renders the home page
func Home(w http.ResponseWriter, r *http.Request) {
	err := renderjet.RenderPage(w, "home.jet", nil)
	if err != nil {
		log.Println(err)
	}
}


// WsHandler estructura para manejar WebSocket
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

// HomePageHandler maneja la solicitud de la página de inicio
func (h *WsHandler) Home(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
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
