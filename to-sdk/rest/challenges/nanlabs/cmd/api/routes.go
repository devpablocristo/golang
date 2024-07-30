package main

import (
	"net/http"

	"github.com/go-chi/chi"

	chihandler "github.com/devpablocristo/nanlabs/infrastructure/driver-adapter/handler/chi"
)

func SetupRouter(handler *chihandler.ChiHandler) http.Handler {
	r := chi.NewRouter()

	r.Post("/", handler.Task)

	return r
}
