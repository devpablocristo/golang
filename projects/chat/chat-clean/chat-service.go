package chatsrv

import (
	"log"
	"net/http"
	"sync"

	router "github.com/devpablocristo/interviews/b6/chat/infrastructure/router"
	usecases "github.com/devpablocristo/interviews/b6/chat/usecases"
)

func ChatService(wg *sync.WaitGroup) {
	defer wg.Done()
	mux := router.Routes()

	log.Println("Starting channel listener")
	go usecases.ListenToWsChannel()

	port := ":9999"

	log.Println("Chat Service starting server on port:", port)

	_ = http.ListenAndServe(port, mux)
}
