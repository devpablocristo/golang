package shared

import (
	"net/http"
)

// ApiError representa un error en la API con un código de estado y un mensaje.
type ApiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Method  string `json:"method,omitempty"`
}

// NewApiError crea una nueva instancia de ApiError.
func NewApiError(code int, message string) *ApiError {
	return &ApiError{
		Code:    code,
		Message: message,
	}
}

// ApiErrors contiene errores comunes de la API.
// Utilizamos un mapa para facilitar la extensión y mantenimiento de errores.
var ApiErrors = map[string]*ApiError{
	"InvalidJSON":    NewApiError(http.StatusBadRequest, "Invalid JSON"),
	"InternalServer": NewApiError(http.StatusInternalServerError, "Internal server error"),
	"NotFound":       NewApiError(http.StatusNotFound, "Resource not found"),
	"Unauthorized":   NewApiError(http.StatusUnauthorized, "Unauthorized access"),
	"Forbidden":      NewApiError(http.StatusForbidden, "Forbidden access"),
	"BadRequest":     NewApiError(http.StatusBadRequest, "Bad request"),
}
