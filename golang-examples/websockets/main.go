package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func reader(ws *websocket.Conn) {
	for {
		// read in a message
		msgType, msg, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		// print out that message for clarity
		log.Println(string(msg))

		// send it back to the client
		if err := ws.WriteMessage(msgType, msg); err != nil {
			log.Println(err)
			return
		}
	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	// accept any conection from any client
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// upgrade the connection to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	// listen for messages to come through on the websocket
	for {
		// read in a message
		msgType, msg, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		// print out that message for clarity
		fmt.Println(string(msg))

		// send it back to the client
		if err := ws.WriteMessage(msgType, msg); err != nil {
			log.Println(err)
			return
		}
	}

}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}

func main() {
	fmt.Println("Go Websockets!")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
