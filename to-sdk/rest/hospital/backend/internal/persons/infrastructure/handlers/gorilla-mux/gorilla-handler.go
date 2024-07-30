package goriadapter

import (
	"github.com/gorilla/mux"
)

type GorillaHandler struct {
	service   ports.PersonService
	muxRouter *mux.Router
}

func NewGorillaHandler(ps ports.PersonService, mr *mux.Router) *GorillaHandler {
	//r := mux.NewRouter()

	return &GorillaHandler{
		service:   ps,
		muxRouter: mr,
	}
}

func (h *GorillaHandler) SetupRoutes() {
	h.muxRouter.HandleFunc("/person", h.List).Methods("GET")
	h.muxRouter.HandleFunc("/person", h.Register).Methods("POST")
}
