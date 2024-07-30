package api

import (
	"github.com/go-chi/chi"

	handler "github.com/devpablocristo/blankfactor/event-service/internal/infrastructure/driver-adapter/handler"
)

func Router(handler *handler.Handler) *chi.Mux {
	router := chi.NewRouter()
	router.Route("/api/v1", func(r chi.Router) {
		r.Route("/events", func(r chi.Router) {
			r.Get("/get-overlaping", handler.GetOverlappingEvents)
			r.Get("/get-all", handler.GetAllEvents)
			r.Post("/create", handler.CreateEvent)
		})
	})
	return router
}
