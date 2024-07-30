package api

import (
	handler "github.com/devpablocristo/interviews/bookstore/src/inventory/adapters"
	"github.com/gorilla/mux"
)

func routes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/inventory", handler.GetBook().Methods("GET")
	return router
}
