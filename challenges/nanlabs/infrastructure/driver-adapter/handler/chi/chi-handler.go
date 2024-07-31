package chihandler

import (
	"encoding/json"
	"log"
	"net/http"

	port "github.com/devpablocristo/nanlabs/application/port"
	domain "github.com/devpablocristo/nanlabs/domain"
	cdomain "github.com/devpablocristo/nanlabs/internal/commons/domain"
)

type ChiHandler struct {
	taskService port.Service
}

func NewChiHandler(ts port.Service) *ChiHandler {
	return &ChiHandler{
		taskService: ts,
	}
}

func (h *ChiHandler) Task(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//w.Write([]byte("hola"))

	body := r.Body
	defer body.Close()

	var errReq cdomain.APIError
	errReq.Method = "chihandler.Task"

	var newTask *domain.Task
	err := json.NewDecoder(body).Decode(&newTask)
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
			//w.Write([]byte(errReq.Message + " - " + errReq.Error))
			return
		}
		log.Println(errReq)
		//w.Write([]byte(errReq.Message + " - " + errReq.Error))
		return
	}

	ctx := r.Context()
	err = h.taskService.CreateCard(ctx, newTask)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		errReq = cdomain.ErrUnauthorized
		errReq.Error = err.Error()
		err := json.NewEncoder(w).Encode(
			cdomain.ErrInternalServer,
		)
		if err != nil {
			errReq.Error = err.Error()
			log.Println(errReq)
			//w.Write([]byte(errReq.Message + " - " + errReq.Error))
			return
		}
		log.Println(errReq)
		//w.Write([]byte(errReq.Message + " - " + errReq.Error))
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(
		cdomain.ResponseAPI{
			Success: true,
			Status:  http.StatusCreated,
			Result:  "Trello card created",
		},
	)
}
