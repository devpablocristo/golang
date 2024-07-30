package api

import (
	"github.com/go-chi/chi"

	handler "github.com/devpablocristo/99minutos/order-manager/internal/infrastructure/driver-adapter/handler"
)

func Router(handler *handler.Handler) *chi.Mux {
	router := chi.NewRouter()
	router.Route("/api/v1", func(r chi.Router) {
		r.Route("/orders", func(r chi.Router) {
			r.Post("/create", handler.CreateOrder)
		})
	})
	return router
}
