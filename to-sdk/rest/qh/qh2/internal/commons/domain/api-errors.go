package cmsdomain

import (
	"encoding/json"
	"net/http"
)

var (
	ErrBadRequest = APIError{
		StatusCode: http.StatusBadRequest,
		Type:       "client-error",
		Error:      "Default",
		Message:    "Cannot process current request",
		Method:     "Default",
	}
	ErrInvalidJSON = APIError{
		StatusCode: http.StatusBadRequest,
		Type:       "invalid-json",
		Error:      "Default",
		Message:    "Invalid or malformed JSON",
		Method:     "Default",
	}
	ErrInternalServer = APIError{
		StatusCode: http.StatusInternalServerError,
		Type:       "server-error",
		Error:      "Default",
		Message:    "Internal server error",
		Method:     "Default",
	}
)

type APIError struct {
	StatusCode int    `json:"-"`
	Type       string `json:"type"`
	Error      string `json:"error"`
	Message    string `json:"message,omitempty"`
	Method     string `json:"method,omitempty"`
}

func (a APIError) Send(w http.ResponseWriter) error {
	statusCode := a.StatusCode
	if statusCode == 0 {
		statusCode = http.StatusBadRequest
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(a)
}

// integrar
type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

// integrar
func BadRequestError(msg string) *RestErr {
	return &RestErr{
		Message: msg,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}
