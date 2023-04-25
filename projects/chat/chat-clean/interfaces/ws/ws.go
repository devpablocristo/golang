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
