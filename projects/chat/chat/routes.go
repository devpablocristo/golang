package chatsrv

import (
	"net/http"

	"github.com/bmizerany/pat"

	handlers "github.com/devpablocristo/interviews/b6/chat/handlers"
)

// routes defines the application routes
func routes() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Home))
	mux.Get("/ws", http.HandlerFunc(handlers.WsEndpoint))

	return mux
}
