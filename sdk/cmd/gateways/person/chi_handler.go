package person

import (
	"encoding/json"
	"net/http"

	"github.com/devpablocristo/golang/sdk/cmd/gateways/shared"
	"github.com/devpablocristo/golang/sdk/internal/core"
)

type ChiHandler struct {
	uc core.PersonUseCasesPort
}

func NewChiHandler(uc core.PersonUseCasesPort) *ChiHandler {
	return &ChiHandler{
		uc: uc,
	}
}

func (h *ChiHandler) CreatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()

	var dto PersonDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		shared.WriteErrorResponse(w, shared.ApiErrors["InvalidJSON"], "ChiHandler.CreatePerson")
		return
	}

	if err := h.uc.CreatePerson(r.Context(), dto.ToDomain()); err != nil {
		shared.WriteErrorResponse(w, shared.ApiErrors["InternalServer"], "ChiHandler.CreatePerson")
		return
	}

	response := shared.NewApiResponse(true, http.StatusCreated, "Person created successfully", dto.ToDomain())
	shared.WriteJSONResponse(w, http.StatusCreated, response)
}

// func (h *ChiHandler) GetPerson(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	UUID := chi.URLParam(r, "personUUID")
// 	var errReq shared.ApiError
// 	errReq.Method = "chihandler.GetPerson"

// 	ctx := r.Context()
// 	person, err := h.uc.GetPerson(ctx, UUID)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		errReq = shared.ErrInternalServer
// 		errReq.Error = err.Error()
// 		err := json.NewEncoder(w).Encode(
// 			shared.ErrInternalServer,
// 		)
// 		if err != nil {
// 			errReq.Error = err.Error()
// 			log.Println(errReq)
// 			w.Write([]byte(errReq.Message + " - " + errReq.Error))
// 			return
// 		}
// 		log.Println(errReq)
// 		w.Write([]byte(errReq.Message + " - " + errReq.Error))
// 		return
// 	}

// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(
// 		shared.ResponseAPI{
// 			Success: true,
// 			Status:  http.StatusCreated,
// 			Result:  person,
// 		},
// 	)
// }

// func (h *ChiHandler) GetPersons(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	var errReq shared.ApiError
// 	errReq.Method = "chihandler.GetPersons"

// 	ctx := r.Context()
// 	persons, err := h.uc.GetPersons(ctx)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		errReq = shared.ErrInternalServer
// 		errReq.Error = err.Error()
// 		err := json.NewEncoder(w).Encode(
// 			shared.ErrInternalServer,
// 		)
// 		if err != nil {
// 			errReq.Error = err.Error()
// 			log.Println(errReq)
// 			w.Write([]byte(errReq.Message + " - " + errReq.Error))
// 			return
// 		}
// 		log.Println(errReq)
// 		w.Write([]byte(errReq.Message + " - " + errReq.Error))
// 		return
// 	}

// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(
// 		shared.ResponseAPI{
// 			Success: true,
// 			Status:  http.StatusCreated,
// 			Result:  persons,
// 		},
// 	)
// }

// func (h *ChiHandler) UpdatePerson(w http.ResponseWriter, r *http.Request) {}
// func (h *ChiHandler) DeletePerson(w http.ResponseWriter, r *http.Request) {}
