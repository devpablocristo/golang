package handler

import (
	"encoding/json"
	"log"
	"net/http"

	cdomain "github.com/devpablocristo/golang/06-projects/qh/internal/commons/domain"
	port "github.com/devpablocristo/golang/06-projects/qh/person/application/port"
	domain "github.com/devpablocristo/golang/06-projects/qh/person/domain"
	"github.com/go-chi/chi"
)

type ChiHandler struct {
	personService port.Service
}

func NewChiHandler(ps port.Service) *ChiHandler {
	return &ChiHandler{
		personService: ps,
	}
}

func (h *ChiHandler) CreatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body := r.Body
	defer body.Close()

	var errReq cdomain.APIError
	errReq.Method = "chiAdapter.CreatePerson"

	var newPerson domain.Person
	err := json.NewDecoder(body).Decode(&newPerson)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errReq = cdomain.ErrInvalidJSON
		errReq.Error = err.Error()
		err := json.NewEncoder(w).Encode(
			errReq,
		)
		if err != nil {
			errReq.Error = err.Error()
			log.Println(errReq)
			w.Write([]byte(errReq.Message + " - " + errReq.Error))
			return
		}
		log.Println(errReq)
		w.Write([]byte(errReq.Message + " - " + errReq.Error))
		return
	}

	ctx := r.Context()
	person, err := h.personService.CreatePerson(ctx, &newPerson)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errReq = cdomain.ErrInternalServer
		errReq.Error = err.Error()
		err := json.NewEncoder(w).Encode(
			cdomain.ErrInternalServer,
		)
		if err != nil {
			errReq.Error = err.Error()
			log.Println(errReq)
			w.Write([]byte(errReq.Message + " - " + errReq.Error))
			return
		}
		log.Println(errReq)
		w.Write([]byte(errReq.Message + " - " + errReq.Error))
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(
		cdomain.ResponseAPI{
			Success: true,
			Status:  http.StatusCreated,
			Result:  person,
		},
	)
}

func (h *ChiHandler) GetPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	UUID := chi.URLParam(r, "personUUID")
	var errReq cdomain.APIError
	errReq.Method = "chihandler.GetPerson"

	ctx := r.Context()
	person, err := h.personService.GetPerson(ctx, UUID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errReq = cdomain.ErrInternalServer
		errReq.Error = err.Error()
		err := json.NewEncoder(w).Encode(
			cdomain.ErrInternalServer,
		)
		if err != nil {
			errReq.Error = err.Error()
			log.Println(errReq)
			w.Write([]byte(errReq.Message + " - " + errReq.Error))
			return
		}
		log.Println(errReq)
		w.Write([]byte(errReq.Message + " - " + errReq.Error))
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(
		cdomain.ResponseAPI{
			Success: true,
			Status:  http.StatusCreated,
			Result:  person,
		},
	)
}

func (h *ChiHandler) GetPersons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var errReq cdomain.APIError
	errReq.Method = "chihandler.GetPersons"

	ctx := r.Context()
	persons, err := h.personService.GetPersons(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errReq = cdomain.ErrInternalServer
		errReq.Error = err.Error()
		err := json.NewEncoder(w).Encode(
			cdomain.ErrInternalServer,
		)
		if err != nil {
			errReq.Error = err.Error()
			log.Println(errReq)
			w.Write([]byte(errReq.Message + " - " + errReq.Error))
			return
		}
		log.Println(errReq)
		w.Write([]byte(errReq.Message + " - " + errReq.Error))
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(
		cdomain.ResponseAPI{
			Success: true,
			Status:  http.StatusCreated,
			Result:  persons,
		},
	)
}

func (h *ChiHandler) UpdatePerson(w http.ResponseWriter, r *http.Request) {}
func (h *ChiHandler) DeletePerson(w http.ResponseWriter, r *http.Request) {}
