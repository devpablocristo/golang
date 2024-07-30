package muxrouter

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	muxRouterInstance = mux.NewRouter()
)

type MuxRouterRepository interface {
	GET(path string, handler func(w http.ResponseWriter, r *http.Request))
	POST(path string, handler func(w http.ResponseWriter, r *http.Request))
	PUT(path string, handler func(w http.ResponseWriter, r *http.Request))
	PATCH(path string, handler func(w http.ResponseWriter, r *http.Request))
	DELETE(path string, handler func(w http.ResponseWriter, r *http.Request))
}

type MuxRouter struct {
}

func NewMuxRouter() *MuxRouter {
	return &MuxRouter{}
}

func (mx MuxRouter) GET(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	muxRouterInstance.HandleFunc(path, handler).Methods("GET")
}

func (mx MuxRouter) POST(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	muxRouterInstance.HandleFunc(path, handler).Methods("POST")
}

func (mx MuxRouter) PUT(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	muxRouterInstance.HandleFunc(path, handler).Methods("PUT")
}

func (mx MuxRouter) PATCH(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	muxRouterInstance.HandleFunc(path, handler).Methods("PATCH")
}

func (mx MuxRouter) DELETE(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	muxRouterInstance.HandleFunc(path, handler).Methods("DELETE")
}

func (mx MuxRouter) SERVE(port string) {
	log.Println("Inventory Service starting server on port:", port)
	http.ListenAndServe(port, muxRouterInstance)
}
