package shared

import (
	"encoding/json"
	"log"
	"net/http"
)

// ApiResponse representa una respuesta de API con un mensaje opcional y un resultado.
type ApiResponse struct {
	Success bool   `json:"success"`
	Status  int    `json:"status"`
	Message string `json:"message,omitempty"`
	Result  any    `json:"result,omitempty"`
}

// NewApiResponse crea una nueva instancia de ApiResponse.
func NewApiResponse(success bool, status int, message string, result any) *ApiResponse {
	return &ApiResponse{
		Success: success,
		Status:  status,
		Message: message,
		Result:  result,
	}
}

// WriteJSONResponse escribe una respuesta JSON al cliente.
func WriteJSONResponse(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Error al codificar la respuesta JSON: %v", err)
	}
}

// WriteErrorResponse escribe una respuesta de error JSON al cliente.
func WriteErrorResponse(w http.ResponseWriter, apiErr *ApiError, method string) {
	apiErr.Method = method // Asigna el método al error.
	response := NewApiResponse(false, apiErr.Code, apiErr.Message, nil)
	WriteJSONResponse(w, apiErr.Code, response)
}
