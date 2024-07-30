package api

import (
	"github.com/go-chi/chi"

	handler "github.com/devpablocristo/growuphr/number-manager/infrastructure/driver-adapter/handler"
)

func Router(handler *handler.Handler) *chi.Mux {
	router := chi.NewRouter()
	router.Route("/api/v1", func(r chi.Router) {
		r.Route("/number-manager", func(r chi.Router) {
			r.Post("/reserve", handler.ReserveNumber)
			r.Post("/reserve/{username}", handler.ReserveNumber)
			r.Get("/reserved-numbers", handler.ReservedNumbers)
		})
	})
	return router
}
