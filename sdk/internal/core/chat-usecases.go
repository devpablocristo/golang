package core

import (
	"fmt"
	"log"
	"net/http"
	"sort"
	"sync"

	"github.com/devpablocristo/golang/sdk/cmd/websocket/chat/router"
	"github.com/devpablocristo/golang/sdk/internal/core/chat"
)

var wsChan = make(chan chat.WsPayload)

var Clients = make(map[chat.WebSocketConnection]string)

func ListenForWs(conn *chat.WebSocketConnection) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Error", fmt.Sprintf("%v", r))
		}
	}()

	var payload chat.WsPayload

	for {
		err := conn.ReadJSON(&payload)
		if err != nil {
			// do nothing
		} else {
			payload.Conn = *conn
			wsChan <- payload
		}
	}
}

func ListenToWsChannel() {
	var response chat.WsJsonResponse

	for {
		e := <-wsChan

		switch e.Action {
		case "username":
			// get a list of all users and send it back via broadcast
			Clients[e.Conn] = e.Username
			users := getUserList()
			response.Action = "list_users"
			response.ConnectedUsers = users
			broadcastToAll(response)

		case "left":
			// handle the situation where a user leaves the page
			response.Action = "list_users"
			delete(Clients, e.Conn)
			users := getUserList()
			response.ConnectedUsers = users
			broadcastToAll(response)

		}
	}
}

func getUserList() []string {
	var userList []string
	for _, x := range Clients {
		if x != "" {
			userList = append(userList, x)
		}
	}
	sort.Strings(userList)
	return userList
}

func broadcastToAll(response chat.WsJsonResponse) {
	for client := range Clients {
		err := client.WriteJSON(response)
		if err != nil {
			log.Println("websocket err")
			_ = client.Close()
			delete(Clients, client)
		}
	}
}

func ChatService(wg *sync.WaitGroup) {
	defer wg.Done()
	mux := router.Routes()

	log.Println("Starting channel listener")
	go usecases.ListenToWsChannel()

	port := ":9999"

	log.Println("Chat Service starting server on port:", port)

	_ = http.ListenAndServe(port, mux)
}
