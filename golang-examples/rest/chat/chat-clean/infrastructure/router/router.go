package router

import (
	"net/http"

	"github.com/bmizerany/pat"
	ws "github.com/devpablocristo/interviews/b6/chat/interfaces/ws"
)

// routes defines the application routes

func Routes() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(ws.Home))
	mux.Get("/ws", http.HandlerFunc(ws.WsEndpoint))

	return mux
}
