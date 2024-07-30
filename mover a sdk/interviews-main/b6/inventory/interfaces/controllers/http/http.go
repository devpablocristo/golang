package http

import (
	"encoding/json"
	"net/http"

	inventory "github.com/devpablocristo/interviews/b6/inventory/domain"
	usecases "github.com/devpablocristo/interviews/b6/inventory/usecases"
)

type HTTPInteractor struct {
	handler usecases.UseCasesInteractor
}

type ErrorResponse struct {
	Message string `json:"error"`
}

func NewHTTPInteractor(handler usecases.UseCasesInteractor) *HTTPInteractor {
	return &HTTPInteractor{handler}
}

func MakeHTTPInteractor(handler usecases.UseCasesInteractor) HTTPInteractor {
	return HTTPInteractor{handler}
}

func (h HTTPInteractor) Add(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var book inventory.Book

	err := json.NewDecoder(req.Body).Decode(&book)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: "Invalid Payload"})
		return
	}

	err2 := h.handler.SaveBook(book)
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: err2.Error()})
		return
	}

	res.WriteHeader(http.StatusOK)
}

func (h HTTPInteractor) GetAll(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	results, err := h.handler.ListInventory()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: err.Error()})
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(results)
}
