package chatsrv

import (
	"log"
	"net/http"
	"sync"

	handlers "github.com/devpablocristo/interviews/b6/chat/handlers"
)

func ChatService(wg *sync.WaitGroup) {
	defer wg.Done()
	mux := routes()

	log.Println("Starting channel listener")
	go handlers.ListenToWsChannel()

	port := ":9999"

	log.Println("Inventory Service starting server on port:", port)

	_ = http.ListenAndServe(port, mux)
}
