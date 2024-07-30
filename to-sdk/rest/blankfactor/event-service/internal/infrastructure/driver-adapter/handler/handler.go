package handler

import (
	"encoding/json"
	"net/http"

	port "github.com/devpablocristo/blankfactor/event-service/internal/application/port"
	domain "github.com/devpablocristo/blankfactor/event-service/internal/domain"
)

type Handler struct {
	eventServ port.EventService
}

func NewHandler(es port.EventService) *Handler {
	return &Handler{
		eventServ: es,
	}
}

func (h *Handler) GetOverlappingEvents(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	events, err := h.eventServ.GetOverlappingEvents(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(events)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func (h *Handler) GetAllEvents(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	events, err := h.eventServ.GetAllEvents(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}

func (h *Handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var data domain.Event
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdEvent, err := h.eventServ.CreateEvent(ctx, &data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdEvent)
}
