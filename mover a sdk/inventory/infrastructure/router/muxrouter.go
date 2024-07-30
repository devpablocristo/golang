package muxrouter

import (
	"net/http"

	"github.com/gorilla/mux"
)

var (
	muxRouterInstance = mux.NewRouter()
)

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
